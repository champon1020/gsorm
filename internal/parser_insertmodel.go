package internal

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/morikuni/failure"
)

// InsertModelParser is the model parser for insert statement.
type InsertModelParser struct {
	Model       reflect.Value
	ModelType   reflect.Type
	Cols        []string
	ColumnField map[int]int
}

// NewInsertModelParser creates InsertModelParser instance.
func NewInsertModelParser(cols []string, model interface{}) (*InsertModelParser, error) {
	mTyp := reflect.TypeOf(model)
	if mTyp.Kind() != reflect.Ptr {
		err := failure.New(errInvalidValue, failure.Message("model must be a pointer"))
		return nil, err
	}

	parser := &InsertModelParser{
		Model:     reflect.ValueOf(model).Elem(),
		ModelType: mTyp.Elem(),
		Cols:      cols,
	}
	return parser, nil
}

// Parse converts model to SQL.
func (p *InsertModelParser) Parse() (*SQL, error) {
	var sql SQL

	switch p.ModelType.Kind() {
	case reflect.Slice, reflect.Array:
		if p.ModelType.Elem().Kind() == reflect.Struct {
			p.ParseStructSlice(&sql, p.Model)
			return &sql, nil
		}
		if p.ModelType.Elem().Kind() == reflect.Map {
			if err := p.ParseMapSlice(&sql, p.Model); err != nil {
				return nil, err
			}
			return &sql, nil
		}
		if err := p.ParseVarSlice(&sql, p.Model); err != nil {
			return nil, err
		}
		return &sql, nil
	case reflect.Struct:
		p.ParseStruct(&sql, p.Model)
		return &sql, nil
	case reflect.Map:
		if err := p.ParseMap(&sql, p.Model); err != nil {
			return nil, err
		}
		return &sql, nil
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Bool,
		reflect.String:
		if err := p.ParseVar(&sql, p.Model); err != nil {
			return nil, err
		}
		return &sql, nil
	}

	err := failure.New(errInvalidType,
		failure.Context{"type": p.ModelType.Kind().String()},
		failure.Message("invalid type for internal.InsertModelParser.Parse"))
	return nil, err
}

// ParseMapSlice parses slice or array of map to SQL.
func (p *InsertModelParser) ParseMapSlice(sql *SQL, model reflect.Value) error {
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
func (p *InsertModelParser) ParseStructSlice(sql *SQL, model reflect.Value) {
	p.ColumnField = p.columnsAndFields(model.Type().Elem())
	for i := 0; i < model.Len(); i++ {
		if i > 0 {
			sql.Write(",")
		}
		p.ParseStruct(sql, model.Index(i))
	}
}

// ParseVarSlice parses slice or array of variable to SQL.
func (p *InsertModelParser) ParseVarSlice(sql *SQL, model reflect.Value) error {
	for i := 0; i < model.Len(); i++ {
		if i > 0 {
			sql.Write(",")
		}
		err := p.ParseVar(sql, model.Index(i))
		if err != nil {
			return err
		}
	}
	return nil
}

// ParseMap parses map to SQL.
func (p *InsertModelParser) ParseMap(sql *SQL, model reflect.Value) error {
	sql.Write("(")
	for i, c := range p.Cols {
		if i > 0 {
			sql.Write(",")
		}
		v := model.MapIndex(reflect.ValueOf(c))
		if !v.IsValid() {
			return failure.New(errInvalidSyntax,
				failure.Message("column names must be included in one of map keys"))
		}
		s := ToString(v.Interface(), nil)
		sql.Write(s)
	}
	sql.Write(")")
	return nil
}

// ParseStruct parses struct to SQL.
func (p *InsertModelParser) ParseStruct(sql *SQL, model reflect.Value) {
	if p.ColumnField == nil {
		p.ColumnField = p.columnsAndFields(model.Type())
	}
	sql.Write("(")
	tags := ExtractTags(reflect.TypeOf(model.Interface()))
	for i := 0; i < len(p.Cols); i++ {
		if i > 0 {
			sql.Write(",")
		}

		var opt *ToStringOpt
		if tags[p.ColumnField[i]].Layout != "" {
			opt = &ToStringOpt{Quotes: true, TimeFormat: tags[p.ColumnField[i]].Layout}
		} else {
			opt = nil
		}
		s := ToString(model.Field(p.ColumnField[i]).Interface(), opt)
		sql.Write(s)
	}
	sql.Write(")")
}

// ParseVar parses variable to SQL.
func (p *InsertModelParser) ParseVar(sql *SQL, model reflect.Value) error {
	if len(p.Cols) != 1 {
		return failure.New(errInvalidSyntax,
			failure.Context{"n_columns": strconv.Itoa(len(p.Cols))},
			failure.Message("invalid number of columns"))
	}

	s := ToString(model.Interface(), nil)
	sql.Write(fmt.Sprintf("(%s)", s))
	return nil
}

func (p *InsertModelParser) columnsAndFields(target reflect.Type) map[int]int {
	cf := make(map[int]int)
	for i, col := range p.Cols {
		for j := 0; j < target.NumField(); j++ {
			c := ExtractTag(target.Field(j)).Column
			if c == "" {
				c = SnakeCase(target.Field(j).Name)
			}
			if col != c {
				continue
			}
			cf[i] = j
		}
	}
	return cf
}
