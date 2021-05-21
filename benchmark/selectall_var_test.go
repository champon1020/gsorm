package benchmark

import (
	"testing"

	"github.com/champon1020/gsorm"
)

func BenchmarkSelectAll_Var_gsorm(b *testing.B) {
	db, err := gsorm.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	var emp []int
	err = gsorm.Select(db, "emp_no").From("employees").Query(&emp)
	if err != nil {
		b.Fatal(err)
	}
}
