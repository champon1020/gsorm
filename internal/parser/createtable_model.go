package parser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/internal"
	"github.com/morikuni/failure"
)

// CreateTableModelParser is the model parser for create table statement.
type CreateTableModelParser struct {
	Model     reflect.Value
	ModelType reflect.Type
	DBDriver  domain.SQLDriver

	f   reflect.StructField
	tag *internal.Tag
	uc  map[string][]string
	pk  map[string][]string
	fk  map[string][]string
	ref map[string]string
}

// NewCreateTableModelParser creates CreateTableModelParser instance.
func NewCreateTableModelParser(model interface{}, driver domain.SQLDriver) (*CreateTableModelParser, error) {
	mt := reflect.TypeOf(model)
	if mt.Kind() != reflect.Ptr {
		err := failure.New(errInvalidValue, failure.Message("model must be a pointer"))
		return nil, err
	}
	mt = mt.Elem()

	m := reflect.ValueOf(model).Elem()

	parser := &CreateTableModelParser{
		Model:     m,
		ModelType: mt,
		DBDriver:  driver,
		uc:        make(map[string][]string),
		pk:        make(map[string][]string),
		fk:        make(map[string][]string),
		ref:       make(map[string]string),
	}

	return parser, nil
}

// Parse converts model to SQL.
func (p *CreateTableModelParser) Parse() (*internal.SQL, error) {
	var sql internal.SQL

	if p.ModelType.Kind() != reflect.Struct {
		err := failure.New(errInvalidValue,
			failure.Context{"type": p.ModelType.Kind().String()},
			failure.Message("invalid type for parser.CreateTableModelParser"))
		return nil, err
	}

	sql.Write("(")
	for i := 0; i < p.ModelType.NumField(); i++ {
		if i > 0 {
			sql.Write(",")
		}

		p.f = p.ModelType.Field(i)
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
func (p *CreateTableModelParser) ParseColumn(sql *internal.SQL) string {
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
func (p *CreateTableModelParser) ParseType(sql *internal.SQL) error {
	if p.tag.Lookup("typ") {
		sql.Write(p.tag.Type)
		return nil
	}

	t := p.DBDriver.LookupDefaultType(p.f.Type)
	if t == "" {
		return failure.New(errInvalidType,
			failure.Context{"type": p.f.Type.String()},
			failure.Message("invalid type for database column"))

	}

	sql.Write(t)
	return nil
}

// ParseNotNull parses the not null property from field tag.
func (p *CreateTableModelParser) ParseNotNull(sql *internal.SQL) {
	if p.tag.Lookup("notnull") {
		sql.Write("NOT NULL")
	}
}

// ParseDefault parses the default property from field tag.
func (p *CreateTableModelParser) ParseDefault(sql *internal.SQL) {
	if p.tag.Lookup("default") {
		sql.Write(fmt.Sprintf("DEFAULT %s", p.tag.Default))
	}
}

// ParseUnique parses the unique key from field tag.
func (p *CreateTableModelParser) ParseUnique(sql *internal.SQL, column string) {
	if p.tag.Lookup("uc") {
		p.uc[p.tag.UC] = append(p.uc[p.tag.UC], column)
	}
}

// ParsePrimary parses the primary key from field tag.
func (p *CreateTableModelParser) ParsePrimary(sql *internal.SQL, column string) {
	if p.tag.Lookup("pk") {
		p.pk[p.tag.PK] = append(p.pk[p.tag.PK], column)
	}
}

// ParseForeign parses the foreign key from field tag.
func (p *CreateTableModelParser) ParseForeign(sql *internal.SQL, column string) {
	if p.tag.Lookup("fk") {
		p.fk[p.tag.FK] = append(p.fk[p.tag.FK], column)
		p.ref[p.tag.FK] = p.tag.Ref
	}
}
