package database_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/champon1020/mgorm/database"
	"gotest.tools/v3/assert"
)

var (
	String  = reflect.TypeOf("")
	Int     = reflect.TypeOf(0)
	Float64 = reflect.TypeOf(0.0)
	Time    = reflect.TypeOf(time.Time{})
	Bool    = reflect.TypeOf(false)
)

func TestSQLDriver_LookupDefaultType(t *testing.T) {
	{
		mysql := database.MysqlDriver
		psql := database.PsqlDriver
		assert.Equal(t, mysql.LookupDefaultType(String), "VARCHAR(128)")
		assert.Equal(t, mysql.LookupDefaultType(Int), "INT")
		assert.Equal(t, mysql.LookupDefaultType(Float64), "FLOAT")
		assert.Equal(t, psql.LookupDefaultType(Float64), "NUMERIC")
		assert.Equal(t, mysql.LookupDefaultType(Time), "DATE")
		assert.Equal(t, mysql.LookupDefaultType(Bool), "SMALLINT")
	}
}
