package dockertest_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
)

func TestGroupBy(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result map[string]int
	}{
		// SELECT title, COUNT(title) FROM titles GROUP BY title;
		{
			mgorm.Select(db, "title", "COUNT(title)").
				From("titles").
				GroupBy("title").(*mgorm.Stmt),
			map[string]int{
				"Engineer":        1,
				"Senior Engineer": 4,
				"Senior Staff":    2,
				"Staff":           3,
			},
		},

		// SELECT title, COUNT(title) FROM titles GROUP BY title HAVING COUNT(title) != 1;
		{
			mgorm.Select(db, "title", "COUNT(title)").
				From("titles").
				GroupBy("title").
				Having("COUNT(title) != ?", 1).(*mgorm.Stmt),
			map[string]int{
				"Senior Engineer": 4,
				"Senior Staff":    2,
				"Staff":           3,
			},
		},
	}

	for i, testCase := range testCases {
		model := make(map[string]int)
		if err := testCase.Stmt.Query(&model); err != nil {
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
