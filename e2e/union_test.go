package e2e_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/statement"
	"github.com/google/go-cmp/cmp"
)

func UnionTest(t *testing.T) {
	testCases := []struct {
		Stmt   *statement.SelectStmt
		Result *[]time.Time
	}{
		// SELECT hire_date AS date FROM employees
		// UNION
		// SELECT from_date AS date FROM salaries
		// LIMIT 5;
		{
			mgorm.Select(db, "hire_date AS date").
				From("employees").
				Union(mgorm.Select(nil, "from_date AS date").
					From("salaries"),
				).
				Limit(5).(*statement.SelectStmt),
			&[]time.Time{
				time.Date(1985, time.February, 16, 0, 0, 0, 0, time.UTC),
				time.Date(1985, time.November, 21, 0, 0, 0, 0, time.UTC),
				time.Date(1986, time.February, 18, 0, 0, 0, 0, time.UTC),
				time.Date(1986, time.June, 26, 0, 0, 0, 0, time.UTC),
				time.Date(1986, time.August, 28, 0, 0, 0, 0, time.UTC),
			},
		},

		// SELECT from_date AS date FROM salaries
		// UNION
		// SELECT from_date AS date FROM titles
		// LIMIT 5;
		{
			mgorm.Select(db, "from_date AS date").
				From("salaries").
				UnionAll(mgorm.Select(nil, "from_date AS date").
					From("titles"),
				).
				Limit(5).(*statement.SelectStmt),
			&[]time.Time{
				time.Date(1985, time.February, 18, 0, 0, 0, 0, time.UTC),
				time.Date(1986, time.February, 18, 0, 0, 0, 0, time.UTC),
				time.Date(1986, time.June, 26, 0, 0, 0, 0, time.UTC),
				time.Date(1986, time.June, 26, 0, 0, 0, 0, time.UTC),
				time.Date(1986, time.December, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]time.Time)
		if err := testCase.Stmt.Query(model); err != nil {
			t.Errorf("Error was occurred: %v", err)
			t.Errorf("Executed SQL: %s", testCase.Stmt.String())
			continue
		}
		if diff := cmp.Diff(testCase.Result, model); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			t.Errorf("Executed SQL: %s", testCase.Stmt.String())
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
