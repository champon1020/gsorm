package benchmark

import (
	"testing"

	"github.com/champon1020/mgorm"
)

func BenchmarkSelectAll_Var_mgorm(b *testing.B) {
	db, err := mgorm.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var emp []int
	err = mgorm.Select(db, "emp_no").From("employees").Query(&emp)
	if err != nil {
		b.Fatal(err)
	}
}
