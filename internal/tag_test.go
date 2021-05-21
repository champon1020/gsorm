package internal_test

import (
	"reflect"
	"testing"

	"github.com/champon1020/gsorm/internal"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

type TagModel struct {
	A string `gsorm:"col,typ=VARCHAR(64),notnull=t,default='test',pk=PK_a,fk=FK_a:reftbl(refcol),uc=UC_a"`
	B string `gsorm:"col" json:"col2"`
}

func TestTag_Lookup(t *testing.T) {
	tag := &internal.Tag{
		Column:  "col",
		Type:    "VARCHAR(64)",
		NotNull: true,
		Default: "'test'",
		PK:      "PK_a",
		FK:      "FK_a",
		Ref:     "reftbl(refcol)",
		UC:      "UC_a",
	}
	assert.Equal(t, true, tag.Lookup("col"))
	assert.Equal(t, true, tag.Lookup("typ"))
	assert.Equal(t, true, tag.Lookup("notnull"))
	assert.Equal(t, true, tag.Lookup("default"))
	assert.Equal(t, true, tag.Lookup("pk"))
	assert.Equal(t, true, tag.Lookup("fk"))
	assert.Equal(t, true, tag.Lookup("uc"))
	assert.Equal(t, false, tag.Lookup("hoge"))
}

func Test_ExtractTags(t *testing.T) {
	expected := []*internal.Tag{
		{
			Column:  "col",
			Type:    "VARCHAR(64)",
			NotNull: true,
			Default: "'test'",
			PK:      "PK_a",
			FK:      "FK_a",
			Ref:     "reftbl(refcol)",
			UC:      "UC_a",
		},
		{Column: "col"},
	}

	tags := internal.ExtractTags(reflect.TypeOf(TagModel{}))
	assert.Equal(t, expected, tags)
}

func Test_ExtractTag(t *testing.T) {
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
		typ := reflect.TypeOf(TagModel{})
		actual := internal.ExtractTag(typ.Field(testCase.FieldNum))
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
