package benchmark

import (
	"testing"

	"github.com/champon1020/mgorm"
)

func BenchmarkSelectAll_Map_mgorm(b *testing.B) {
	db, err := mgorm.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var emp []map[string]interface{}
	err = mgorm.Select(db).From("employees").Query(&emp)
	if err != nil {
		b.Fatal(err)
	}
}
