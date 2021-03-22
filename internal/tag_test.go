package internal_test

import (
	"reflect"
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
)

func Test_ExtractTag(t *testing.T) {
	type Model struct {
		A string `mgorm:"col,typ=VARCHAR(64),notnull=t,default='test',pk=PK_a,fk=FK_a:reftbl(refcol),uc=UC_a,layout=time.RFC3339"`
		B string `mgorm:"col" json:"col2"`
	}

	testCases := []struct {
		FieldNum int
		Expected *internal.Tag
	}{
		{
			0,
			&internal.Tag{
				Column:  "col",
				Type:    "VARCHAR(64)",
				NotNull: true,
				Default: "'test'",
				PK:      "PK_a",
				FK:      "FK_a",
				Ref:     "reftbl(refcol)",
				UC:      "UC_a",
				Layout:  "time.RFC3339",
			},
		},
		{
			1,
			&internal.Tag{
				Column: "col",
			},
		},
	}

	for _, testCase := range testCases {
		typ := reflect.TypeOf(Model{})
		actual := internal.ExtractTag(typ.Field(testCase.FieldNum))
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
