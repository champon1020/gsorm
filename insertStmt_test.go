package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/stretchr/testify/assert"
)

func TestInsertStmt_String(t *testing.T) {
	type Model struct {
		ID   int
		Name string `mgorm:"first_name"`
	}

	model1 := Model{ID: 10000, Name: "Taro"}
	model2 := []Model{{ID: 10000, Name: "Taro"}, {ID: 10001, Name: "Hanako"}}
	model3 := []int{10000, 10001}

	testCases := []struct {
		Stmt     *mgorm.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "sample", "name AS first_name", "id").Model(&model1).(*mgorm.InsertStmt),
			`INSERT INTO sample (name AS first_name, id) VALUES ("Taro", 10000)`,
		},
		{
			mgorm.Insert(nil, "sample", "name AS first_name", "id").Model(&model2).(*mgorm.InsertStmt),
			`INSERT INTO sample (name AS first_name, id) VALUES ("Taro", 10000), ("Hanako", 10001)`,
		},
		{
			mgorm.Insert(nil, "sample", "id").Model(&model3).(*mgorm.InsertStmt),
			`INSERT INTO sample (id) VALUES (10000), (10001)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
