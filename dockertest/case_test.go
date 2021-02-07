package dockertest_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *[]string
	}{
		// SELECT CASE
		// WHEN gender = "M" THEN first_name
		// ELSE last_name
		// END
		// FROM employees ORDER BY emp_no;
		{
			mgorm.Select(db, mgorm.When("gender = ?", "M").
				Then("first_name").
				Else("last_name").Column()).
				From("employees").
				OrderBy("emp_no").(*mgorm.Stmt),
			&[]string{
				"Georgi",
				"Simmel",
				"Parto",
				"Chirstian",
				"Kyoichi",
				"Preusig",
				"Zielinski",
				"Saniya",
				"Peac",
				"Piveteau",
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]string)
		if err := testCase.Stmt.Query(model); err != nil {
			t.Errorf("Error was occurred: %v", err)
			t.Errorf("Executed SQL: %s", testCase.Stmt.String())
			continue
		}
		if diff := cmp.Diff(testCase.Result, model); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			t.Errorf("Executed SQL: %s", testCase.Stmt.String())
			internal.PrintTestDiff(t, diff)
		}
	}
}
