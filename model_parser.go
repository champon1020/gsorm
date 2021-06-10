package gsorm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/champon1020/gsorm/internal"
	"golang.org/x/xerrors"
)

// createTableModelParser is the model parser for create table statement.
type createTableModelParser struct {
	model     reflect.Value
	modelType reflect.Type

	f   reflect.StructField
	tag *internal.Tag
	uc  map[string][]string
	pk  map[string][]string
	fk  map[string][]string
	ref map[string]string
}

// newCreateTableModelParser creates createTableModelParser instance.
func newCreateTableModelParser(model interface{}) (*createTableModelParser, error) {
	mt := reflect.TypeOf(model)
	if mt.Kind() != reflect.Ptr {
		return nil, xerrors.New("model must be a pointer")
	}
	mt = mt.Elem()

	m := reflect.ValueOf(model).Elem()

	parser := &createTableModelParser{
		model:     m,
		modelType: mt,
		uc:        make(map[string][]string),
		pk:        make(map[string][]string),
		fk:        make(map[string][]string),
		ref:       make(map[string]string),
	}

	return parser, nil
}

// Parse converts model to SQL.
func (p *createTableModelParser) Parse() (*internal.SQL, error) {
	var sql internal.SQL

	if p.modelType.Kind() != reflect.Struct {
		return nil, xerrors.Errorf("%s is invalid type for createTableModelParser.Parse", p.modelType.Kind().String())
	}

	sql.Write("(")
	for i := 0; i < p.modelType.NumField(); i++ {
		if i > 0 {
			sql.Write(",")
		}

		p.f = p.modelType.Field(i)
		p.tag = internal.ExtractTag(p.f)

		column := p.ParseColumn(&sql)

		if err := p.ParseType(&sql); err != nil {
			return nil, err
		}

		p.ParseNotNull(&sql)
		p.ParseDefault(&sql)
		p.ParseUnique(&sql, column)
		p.ParsePrimary(&sql, column)
		p.ParseForeign(&sql, column)
	}

	// Write unique key if exist.
	for k, v := range p.uc {
		sql.Write(",")
		sql.Write(fmt.Sprintf("CONSTRAINT %s UNIQUE (%s)", k, strings.Join(v, ", ")))
	}

	// Write primary key if exist.
	for k, v := range p.pk {
		sql.Write(",")
		sql.Write(fmt.Sprintf("CONSTRAINT %s PRIMARY KEY (%s)", k, strings.Join(v, ", ")))
	}

	// Write foreign key if exist.
	for k, v := range p.fk {
		sql.Write(",")
		sql.Write(fmt.Sprintf("CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s", k, strings.Join(v, ", "), p.ref[k]))
	}

	sql.Write(")")

	return &sql, nil
}

// ParseColumn parses the database column name from field tag.
func (p *createTableModelParser) ParseColumn(sql *internal.SQL) string {
	var c string
	if p.tag.Lookup("col") {
		c = p.tag.Column
		sql.Write(c)
		return c
	}
	c = internal.SnakeCase(p.f.Name)
	sql.Write(c)
	return c
}

// ParseType parses the database column type from field tag.
func (p *createTableModelParser) ParseType(sql *internal.SQL) error {
	if p.tag.Lookup("typ") {
		sql.Write(p.tag.Type)
		return nil
	}

	return xerrors.New("typ is required")
}

// ParseNotNull parses the not null property from field tag.
func (p *createTableModelParser) ParseNotNull(sql *internal.SQL) {
	if p.tag.Lookup("notnull") {
		sql.Write("NOT NULL")
	}
}

// ParseDefault parses the default property from field tag.
func (p *createTableModelParser) ParseDefault(sql *internal.SQL) {
	if p.tag.Lookup("default") {
		sql.Write(fmt.Sprintf("DEFAULT %s", p.tag.Default))
	}
}

// ParseUnique parses the unique key from field tag.
func (p *createTableModelParser) ParseUnique(sql *internal.SQL, column string) {
	if p.tag.Lookup("uc") {
		p.uc[p.tag.UC] = append(p.uc[p.tag.UC], column)
	}
}

// ParsePrimary parses the primary key from field tag.
func (p *createTableModelParser) ParsePrimary(sql *internal.SQL, column string) {
	if p.tag.Lookup("pk") {
		p.pk[p.tag.PK] = append(p.pk[p.tag.PK], column)
	}
}

// ParseForeign parses the foreign key from field tag.
func (p *createTableModelParser) ParseForeign(sql *internal.SQL, column string) {
	if p.tag.Lookup("fk") {
		p.fk[p.tag.FK] = append(p.fk[p.tag.FK], column)
		p.ref[p.tag.FK] = p.tag.Ref
	}
}

// insertModelParser is the model parser for insert statement.
type insertModelParser struct {
	model       reflect.Value
	modelType   reflect.Type
	Cols        []string
	ColumnField map[int]int
}

// newInsertModelParser creates insertModelParser instance.
func newInsertModelParser(cols []string, model interface{}) (*insertModelParser, error) {
	mTyp := reflect.TypeOf(model)
	if mTyp.Kind() != reflect.Ptr {
		return nil, xerrors.New("model must be a pointer")
	}

	parser := &insertModelParser{
		model:     reflect.ValueOf(model).Elem(),
		modelType: mTyp.Elem(),
		Cols:      cols,
	}
	return parser, nil
}

// Parse converts model to SQL.
func (p *insertModelParser) Parse() (*internal.SQL, error) {
	var sql internal.SQL

	switch p.modelType.Kind() {
	case reflect.Slice, reflect.Array:
		if p.modelType.Elem().Kind() == reflect.Struct {
			p.ParseStructSlice(&sql, p.model)
			return &sql, nil
		}
		if p.modelType.Elem().Kind() == reflect.Map {
			if err := p.ParseMapSlice(&sql, p.model); err != nil {
				return nil, err
			}
			return &sql, nil
		}
	case reflect.Struct:
		p.ParseStruct(&sql, p.model)
		return &sql, nil
	case reflect.Map:
		if err := p.ParseMap(&sql, p.model); err != nil {
			return nil, err
		}
		return &sql, nil
	}

	return nil, xerrors.Errorf("%s is invalide type for insertModelParser.Parse", p.modelType.Kind().String())
}

// ParseMapSlice parses slice or array of map to SQL.
func (p *insertModelParser) ParseMapSlice(sql *internal.SQL, model reflect.Value) error {
	for i := 0; i < model.Len(); i++ {
		if i > 0 {
			sql.Write(",")
		}
		if err := p.ParseMap(sql, model.Index(i)); err != nil {
			return err
		}
	}
	return nil
}

// ParseStructSlice parses slice or array of struct to SQL.
func (p *insertModelParser) ParseStructSlice(sql *internal.SQL, model reflect.Value) {
	p.ColumnField = p.columnsAndFields(model.Type().Elem())
	for i := 0; i < model.Len(); i++ {
		if i > 0 {
			sql.Write(",")
		}
		p.ParseStruct(sql, model.Index(i))
	}
}

// ParseMap parses map to SQL.
func (p *insertModelParser) ParseMap(sql *internal.SQL, model reflect.Value) error {
	sql.Write("(")
	for i, c := range p.Cols {
		if i > 0 {
			sql.Write(",")
		}
		v := model.MapIndex(reflect.ValueOf(c))
		if !v.IsValid() {
			return xerrors.New("column names must be included in one of map keys")
		}
		s := internal.ToString(v.Interface(), nil)
		sql.Write(s)
	}
	sql.Write(")")
	return nil
}

// ParseStruct parses struct to SQL.
func (p *insertModelParser) ParseStruct(sql *internal.SQL, model reflect.Value) {
	if p.ColumnField == nil {
		p.ColumnField = p.columnsAndFields(model.Type())
	}
	sql.Write("(")
	for i := 0; i < len(p.Cols); i++ {
		if i > 0 {
			sql.Write(",")
		}
		opt := &internal.ToStringOpt{Quotes: true}
		s := internal.ToString(model.Field(p.ColumnField[i]).Interface(), opt)
		sql.Write(s)
	}
	sql.Write(")")
}

func (p *insertModelParser) columnsAndFields(target reflect.Type) map[int]int {
	cf := make(map[int]int)
	for i, col := range p.Cols {
		for j := 0; j < target.NumField(); j++ {
			c := internal.ExtractTag(target.Field(j)).Column
			if c == "" {
				c = internal.SnakeCase(target.Field(j).Name)
			}
			if col != c {
				continue
			}
			cf[i] = j
		}
	}
	return cf
}

// updateModelParser is the model parser for update statement.
type updateModelParser struct {
	model       reflect.Value
	modelType   reflect.Type
	Cols        []string
	ColumnField map[int]int
}

// newUpdateModelParser creates updateModelParser instance.
func newUpdateModelParser(cols []string, model interface{}) (*updateModelParser, error) {
	mTyp := reflect.TypeOf(model)
	if mTyp.Kind() != reflect.Ptr {
		return nil, xerrors.New("model must be a pointer")
	}

	parser := &updateModelParser{
		model:     reflect.ValueOf(model).Elem(),
		modelType: mTyp.Elem(),
		Cols:      cols,
	}
	return parser, nil
}

// Parse converts model to SQL.
func (p *updateModelParser) Parse() (*internal.SQL, error) {
	var sql internal.SQL

	switch p.modelType.Kind() {
	case reflect.Struct:
		p.ParseStruct(&sql, p.model)
		return &sql, nil
	case reflect.Map:
		if err := p.ParseMap(&sql, p.model); err != nil {
			return nil, err
		}
		return &sql, nil
	}

	return nil, xerrors.Errorf("%s is invalid type for insertModelParser.Parse", p.modelType.Kind().String())
}

// ParseMap parses map to SQL.
func (p *updateModelParser) ParseMap(sql *internal.SQL, model reflect.Value) error {
	for i, c := range p.Cols {
		if i > 0 {
			sql.Write(",")
		}
		v := model.MapIndex(reflect.ValueOf(c))
		if !v.IsValid() {
			return xerrors.New("column names must be included in one of map keys")
		}
		s := internal.ToString(v.Interface(), nil)
		sql.Write(fmt.Sprintf("%s = %s", c, s))
	}
	return nil
}

// ParseStruct parses struct to SQL.
func (p *updateModelParser) ParseStruct(sql *internal.SQL, model reflect.Value) {
	if p.ColumnField == nil {
		p.ColumnField = p.columnsAndFields(model.Type())
	}
	for i := 0; i < len(p.Cols); i++ {
		if i > 0 {
			sql.Write(",")
		}
		opt := &internal.ToStringOpt{Quotes: true}
		s := internal.ToString(model.Field(p.ColumnField[i]).Interface(), opt)
		sql.Write(fmt.Sprintf("%s = %s", p.Cols[i], s))
	}
}

func (p *updateModelParser) columnsAndFields(target reflect.Type) map[int]int {
	cf := make(map[int]int)
	for i, col := range p.Cols {
		for j := 0; j < target.NumField(); j++ {
			c := internal.ExtractTag(target.Field(j)).Column
			if c == "" {
				c = internal.SnakeCase(target.Field(j).Name)
			}
			if col != c {
				continue
			}
			cf[i] = j
		}
	}
	return cf
}
