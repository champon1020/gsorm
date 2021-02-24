package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/stretchr/testify/assert"
)

func TestInsertStmt_String(t *testing.T) {
	type Model struct {
		ID        int
		FirstName string `mgorm:"name"`
	}

	model1 := Model{ID: 10000, FirstName: "Taro"}
	model2 := []Model{{ID: 10000, FirstName: "Taro"}, {ID: 10001, FirstName: "Jiro"}}
	model3 := []int{10000, 10001}
	model4 := map[string]interface{}{
		"id":   10000,
		"name": "Taro",
	}

	testCases := []struct {
		Stmt     *mgorm.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "sample", "id", "name").Values(10000, "Taro").(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro")`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").
				Values(10000, "Taro").
				Values(10001, "Jiro").(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro"), (10001, "Jiro")`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").
				Values(10000, "Taro").
				Values(10001, "Jiro").
				Values(10002, "Saburo").(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro"), (10001, "Jiro"), (10002, "Saburo")`,
		},
		// Test for (*InsertStmt).Model
		{
			mgorm.Insert(nil, "sample", "id", "name").Model(&model1).(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro")`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").Model(&model2).(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro"), (10001, "Jiro")`,
		},
		{
			mgorm.Insert(nil, "sample", "id").Model(&model3).(*mgorm.InsertStmt),
			`INSERT INTO sample (id) VALUES (10000), (10001)`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").Model(&model4).(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro")`,
		},
		// Test for mapping.
		{
			mgorm.Insert(nil, "sample", "name", "id").Model(&model1).(*mgorm.InsertStmt),
			`INSERT INTO sample (name, id) VALUES ("Taro", 10000)`,
		},
		{
			mgorm.Insert(nil, "sample", "name", "id").Model(&model2).(*mgorm.InsertStmt),
			`INSERT INTO sample (name, id) VALUES ("Taro", 10000), ("Jiro", 10001)`,
		},
		{
			mgorm.Insert(nil, "sample", "name", "id").Model(&model4).(*mgorm.InsertStmt),
			`INSERT INTO sample (name, id) VALUES ("Taro", 10000)`,
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
