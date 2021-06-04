package parser

import (
	"reflect"

	"github.com/champon1020/gsorm/internal"
	"github.com/morikuni/failure"
)

// InsertModelParser is the model parser for insert statement.
type InsertModelParser struct {
	model       reflect.Value
	modelType   reflect.Type
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
		model:     reflect.ValueOf(model).Elem(),
		modelType: mTyp.Elem(),
		Cols:      cols,
	}
	return parser, nil
}

// Parse converts model to SQL.
func (p *InsertModelParser) Parse() (*internal.SQL, error) {
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

	err := failure.New(errInvalidType,
		failure.Context{"type": p.modelType.Kind().String()},
		failure.Message("invalid type for internal.InsertModelParser.Parse"))
	return nil, err
}

// ParseMapSlice parses slice or array of map to SQL.
func (p *InsertModelParser) ParseMapSlice(sql *internal.SQL, model reflect.Value) error {
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
func (p *InsertModelParser) ParseStructSlice(sql *internal.SQL, model reflect.Value) {
	p.ColumnField = p.columnsAndFields(model.Type().Elem())
	for i := 0; i < model.Len(); i++ {
		if i > 0 {
			sql.Write(",")
		}
		p.ParseStruct(sql, model.Index(i))
	}
}

// ParseMap parses map to SQL.
func (p *InsertModelParser) ParseMap(sql *internal.SQL, model reflect.Value) error {
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
		s := internal.ToString(v.Interface(), nil)
		sql.Write(s)
	}
	sql.Write(")")
	return nil
}

// ParseStruct parses struct to SQL.
func (p *InsertModelParser) ParseStruct(sql *internal.SQL, model reflect.Value) {
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

func (p *InsertModelParser) columnsAndFields(target reflect.Type) map[int]int {
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
