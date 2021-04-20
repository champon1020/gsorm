package internal

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
)

// UpdateModelParser is the model parser for update statement.
type UpdateModelParser struct {
	Model       reflect.Value
	ModelType   reflect.Type
	Cols        []string
	ColumnField map[int]int
}

// NewUpdateModelParser creates UpdateModelParser instance.
func NewUpdateModelParser(cols []string, model interface{}) (*UpdateModelParser, error) {
	mTyp := reflect.TypeOf(model)
	if mTyp.Kind() != reflect.Ptr {
		return nil, errors.New("Model must be pointer", errors.InvalidTypeError)
	}

	parser := &UpdateModelParser{
		Model:     reflect.ValueOf(model).Elem(),
		ModelType: mTyp.Elem(),
		Cols:      cols,
	}
	return parser, nil
}

// Parse converts model to SQL.
func (p *UpdateModelParser) Parse() (*SQL, error) {
	var sql SQL

	switch p.ModelType.Kind() {
	case reflect.Struct:
		p.ParseStruct(&sql, p.Model)
		return &sql, nil
	case reflect.Map:
		if err := p.ParseMap(&sql, p.Model); err != nil {
			return nil, err
		}
		return &sql, nil
	}

	msg := fmt.Sprintf("Type %v is not supported", p.ModelType.Kind())
	return nil, errors.New(msg, errors.InvalidTypeError)
}

// ParseMap parses map to SQL.
func (p *UpdateModelParser) ParseMap(sql *SQL, model reflect.Value) error {
	for i, c := range p.Cols {
		if i > 0 {
			sql.Write(",")
		}
		v := model.MapIndex(reflect.ValueOf(c))
		if !v.IsValid() {
			msg := "Column names must be included in one of map keys"
			return errors.New(msg, errors.InvalidSyntaxError)
		}
		s := ToString(v.Interface(), nil)
		sql.Write(fmt.Sprintf("%s = %s", c, s))
	}
	return nil
}

// ParseStruct parses struct to SQL.
func (p *UpdateModelParser) ParseStruct(sql *SQL, model reflect.Value) {
	if p.ColumnField == nil {
		p.ColumnField = p.columnsAndFields(model.Type())
	}
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
		sql.Write(fmt.Sprintf("%s = %s", p.Cols[i], s))
	}
}

func (p *UpdateModelParser) columnsAndFields(target reflect.Type) map[int]int {
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
