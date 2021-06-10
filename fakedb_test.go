package gsorm_test

import (
	"reflect"
	"time"

	"github.com/champon1020/gsorm"
)

type fakeDB struct {
	r gsorm.ExportedIRows
}

func newFakeDB(r gsorm.ExportedIRows) gsorm.DB {
	return &fakeDB{r: r}
}

func (d *fakeDB) Ping() error {
	return nil
}

func (d *fakeDB) Query(query string, args ...interface{}) (gsorm.ExportedIRows, error) {
	return d.r, nil
}

func (d *fakeDB) Exec(query string, args ...interface{}) (gsorm.ExportedIResult, error) {
	return nil, nil
}

func (d *fakeDB) SetConnMaxLifetime(n time.Duration) error {
	return nil
}

func (d *fakeDB) SetMaxIdleConns(n int) error {
	return nil
}

func (d *fakeDB) SetMaxOpenConns(n int) error {
	return nil
}

func (d *fakeDB) Close() error {
	return nil
}

func (d *fakeDB) Begin() (gsorm.Tx, error) {
	return nil, nil
}

type fakeRows struct {
	ct  []gsorm.ExportedIColumnType
	v   [][]interface{}
	itr int
}

func newFakeRows(ct []gsorm.ExportedIColumnType, v [][]interface{}) gsorm.ExportedIRows {
	return &fakeRows{ct: ct, v: v, itr: -1}
}

func (r *fakeRows) Next() bool {
	if r.itr+1 < len(r.v) {
		r.itr++
		return true
	}
	return false
}

func (r *fakeRows) Scan(args ...interface{}) error {
	for i, a := range args {
		v := reflect.ValueOf(r.v[r.itr][i])
		reflect.ValueOf(a).Elem().Set(v)
	}
	return nil
}

func (r *fakeRows) ColumnTypes() ([]gsorm.ExportedIColumnType, error) {
	return r.ct, nil
}

func (r *fakeRows) Close() error {
	return nil
}

type fakeColumnType struct {
	n string
	t reflect.Type
}

func newFakeColumn(n string, t reflect.Type) gsorm.ExportedIColumnType {
	return &fakeColumnType{n: n, t: t}
}

func (c *fakeColumnType) Name() string {
	return c.n
}

func (c *fakeColumnType) ScanType() reflect.Type {
	return c.t
}
