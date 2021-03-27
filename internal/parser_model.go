package internal

import (
	"reflect"

	"github.com/champon1020/mgorm/errors"
)

type ModelParser struct {
	Model       interface{}
	ModelType   reflect.Type
	Cols        []string
	ColumnField map[int]int
}

func NewModelParser(model interface{}) (*ModelParser, error) {
	mTyp := reflect.TypeOf(model)
	if mTyp.Kind() != reflect.Ptr {
		return nil, errors.New("Model must be pointer", errors.InvalidTypeError)
	}

	parser := &ModelParser{
		Model:     model,
		ModelType: mTyp.Elem(),
	}
	return parser, nil
}

func (p *ModelParser) Parse() (*SQL, error) {
	var sql *SQL
	model := reflect.ValueOf(p.Model)

	switch p.ModelType.Kind() {
	case reflect.Slice,
		reflect.Array:
		if p.ModelType.Elem().Kind() == reflect.Struct {
			p.ParseStructSlice(sql, model)
			return sql, nil
		}
		if p.ModelType.Elem().Kind() == reflect.Map {
			if err := p.ParseMapSlice(sql, model); err != nil {
				return nil, err
			}
			return sql, nil
		}
	case reflect.Struct:
		p.ParseStruct(sql, model)
		return sql, nil
	case reflect.Map:
		if err := p.ParseMap(sql, model); err != nil {
			return nil, err
		}
		return sql, nil
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
	}

	return nil, nil
}

func (p *ModelParser) ParseMapSlice(sql *SQL, model reflect.Value) error {
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

func (p *ModelParser) ParseStructSlice(sql *SQL, model reflect.Value) {
	p.ColumnField = p.columnsAndFields(model.Type())
	for i := 0; i < model.Len(); i++ {
		if i > 0 {
			sql.Write(",")
		}
		p.ParseStruct(sql, model.Index(i))
	}
}

func (p *ModelParser) ParseMap(sql *SQL, model reflect.Value) error {
	sql.Write("(")
	for i, c := range p.Cols {
		if i > 0 {
			sql.Write(",")
		}
		v := model.MapIndex(reflect.ValueOf(c))
		if !v.IsValid() {
			msg := "Column names must be included in oneof map keys"
			return errors.New(msg, errors.InvalidSyntaxError)
		}
		s := ToString(v.Interface(), true)
		sql.Write(s)
	}
	sql.Write(")")
	return nil
}

func (p *ModelParser) ParseStruct(sql *SQL, model reflect.Value) {
	if p.ColumnField == nil {
		p.ColumnField = p.columnsAndFields(model.Type())
	}
	sql.Write("(")
	for i := 0; i < len(p.Cols); i++ {
		if i > 0 {
			sql.Write(",")
		}
		s := ToString(model.Field(p.ColumnField[i]).Interface(), true)
		sql.Write(s)
	}
	sql.Write(")")
}

func (p *ModelParser) columnsAndFields(target reflect.Type) map[int]int {
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
