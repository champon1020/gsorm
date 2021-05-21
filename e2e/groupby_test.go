package e2e_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/statement"
	"github.com/google/go-cmp/cmp"
)

func TestGroupBy(t *testing.T) {
	testCases := []struct {
		Stmt   *statement.SelectStmt
		Result []map[string]interface{}
	}{
		// SELECT title, COUNT(title) FROM titles GROUP BY title;
		{
			gsorm.Select(db, "title", "COUNT(title)").
				From("titles").
				GroupBy("title").(*statement.SelectStmt),
			[]map[string]interface{}{
				{
					"title":        "Engineer",
					"COUNT(title)": 1,
				},
				{
					"title":        "Senior Engineer",
					"COUNT(title)": 4,
				},
				{
					"title":        "Senior Staff",
					"COUNT(title)": 2,
				},
				{
					"title":        "Staff",
					"COUNT(title)": 3,
				},
			},
		},

		// SELECT title, COUNT(title) FROM titles GROUP BY title HAVING COUNT(title) != 1;
		{
			gsorm.Select(db, "title", "COUNT(title)").
				From("titles").
				GroupBy("title").
				Having("COUNT(title) != ?", 1).(*statement.SelectStmt),
			[]map[string]interface{}{
				{
					"title":        "Senior Engineer",
					"COUNT(title)": 4,
				},
				{
					"title":        "Senior Staff",
					"COUNT(title)": 2,
				},
				{
					"title":        "Staff",
					"COUNT(title)": 3,
				},
			},
		},
	}

	for i, testCase := range testCases {
		model := make([]map[string]interface{}, 10)
		if err := testCase.Stmt.Query(&model); err != nil {
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
