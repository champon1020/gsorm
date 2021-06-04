package parser

import (
	"reflect"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/morikuni/failure"
)

// MapRowsToModel executes query and sets rows to model structure.
func MapRowsToModel(r domain.Rows, model interface{}) error {
	ct, err := r.ColumnTypes()
	if err != nil {
		return failure.Wrap(err)
	}

	cti := make([]ColumnType, len(ct))
	for i := 0; i < len(ct); i++ {
		cti[i] = ct[i]
	}

	p, err := NewRowsParser(r, cti, model)
	if err != nil {
		return err
	}

	v, err := p.Parse()
	if err != nil {
		return err
	}

	ref := reflect.ValueOf(model).Elem()
	ref.Set(*v)
	return nil
}
