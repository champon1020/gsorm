package e2e_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/statement"
	"github.com/google/go-cmp/cmp"
)

func TestOrderBy(t *testing.T) {
	testCases := []struct {
		Stmt   *statement.SelectStmt
		Result *[]Employee
	}{
		// SELECT * FROM first_name ORDER BY first_name;
		{
			gsorm.Select(db, "first_name").
				From("employees").
				OrderBy("first_name").(*statement.SelectStmt),
			&[]Employee{
				{FirstName: "Anneke"},
				{FirstName: "Bezalel"},
				{FirstName: "Chirstian"},
				{FirstName: "Duangkaew"},
				{FirstName: "Georgi"},
				{FirstName: "Kyoichi"},
				{FirstName: "Parto"},
				{FirstName: "Saniya"},
				{FirstName: "Sumant"},
				{FirstName: "Tzvetan"},
			},
		},

		// SELECT * FROM first_name ORDER BY first_name DESC;
		{
			gsorm.Select(db, "first_name").
				From("employees").
				OrderBy("first_name DESC").(*statement.SelectStmt),
			&[]Employee{
				{FirstName: "Tzvetan"},
				{FirstName: "Sumant"},
				{FirstName: "Saniya"},
				{FirstName: "Parto"},
				{FirstName: "Kyoichi"},
				{FirstName: "Georgi"},
				{FirstName: "Duangkaew"},
				{FirstName: "Chirstian"},
				{FirstName: "Bezalel"},
				{FirstName: "Anneke"},
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]Employee)
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
