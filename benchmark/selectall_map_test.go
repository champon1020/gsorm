package benchmark

import (
	"testing"

	"github.com/champon1020/gsorm"
)

func BenchmarkSelectAll_Map_gsorm(b *testing.B) {
	db, err := gsorm.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	var emp []map[string]interface{}
	err = gsorm.Select(db).From("employees").Query(&emp)
	if err != nil {
		b.Fatal(err)
	}
}
