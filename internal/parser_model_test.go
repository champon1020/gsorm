package internal_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

func TestModelParser_ParseStruct(t *testing.T) {
	testCases := []struct {
		Model    interface{}
		Expected string
	}{
		{},
	}

	for _, testCase := range testCases {
		parser, err := internal.NewModelParser(testCase.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		sql, err := parser.Parse()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, sql.String())
	}
}
