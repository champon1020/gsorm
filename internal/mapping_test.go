package internal_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestValues2Fields(t *testing.T) {
	type Car struct {
		S   string    //0
		I   int       //1
		I8  int8      //2
		I16 int16     //3
		I32 int32     //4
		I64 int64     //5
		U   uint      //6
		U8  uint8     //7
		U16 uint16    //8
		U32 uint32    //9
		U64 uint64    //10
		F32 float32   //11
		F64 float64   //12
		BL  bool      //13
		T   time.Time //14
		T2  time.Time `layout:"2006-01-02"` //15
		T3  time.Time `layout:"time.ANSIC"` //16
	}

	testCases := []struct {
		Type     reflect.Type
		IdxR2M   map[int]int
		Vals     [][]byte
		Expected reflect.Value
	}{
		{
			reflect.TypeOf(Car{}),
			map[int]int{
				0: 0,
				1: 1,
				2: 11,
				3: 13,
				4: 14,
				5: 15,
				6: 16,
			},
			[][]byte{
				[]byte("str"),
				[]byte("10"),
				[]byte("100.2"),
				[]byte("true"),
				[]byte("2020-01-02T03:04:05Z"),
				[]byte("2020-01-02"),
				[]byte("Thu Jan 2 03:04:05 2020"),
			},
			reflect.ValueOf(Car{
				S:   "str",
				I:   10,
				F32: 100.2,
				BL:  true,
				T:   time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC),
				T2:  time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
				T3:  time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC),
			}),
		},
	}

	for _, testCase := range testCases {
		v, err := internal.Values2Fields(testCase.Type, testCase.IdxR2M, testCase.Vals)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected.Interface(), v.Interface())
	}
}

func TestValue2Map(t *testing.T) {
	testCases := []struct {
		Type          reflect.Type
		Key           string
		Value         string
		ExpectedKey   reflect.Value
		ExpectedValue reflect.Value
	}{
		{
			Type:          reflect.TypeOf(map[int]string{}),
			Key:           "10",
			Value:         "str",
			ExpectedKey:   reflect.ValueOf(10),
			ExpectedValue: reflect.ValueOf("str"),
		},
	}

	for _, testCase := range testCases {
		k, v, err := internal.Value2Map(testCase.Type, internal.Str(testCase.Key), internal.Str(testCase.Value))
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.ExpectedKey.Interface(), (*k).Interface()); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
		if diff := cmp.Diff(testCase.ExpectedValue.Interface(), (*v).Interface()); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestValue2Var(t *testing.T) {
	testCases := []struct {
		Type     reflect.Type
		Value    string
		Expected reflect.Value
	}{
		{
			Type:     reflect.TypeOf(""),
			Value:    "str",
			Expected: reflect.ValueOf("str"),
		},
		{
			Type:     reflect.TypeOf(int(0)),
			Value:    "10",
			Expected: reflect.ValueOf(10),
		},
		{
			Type:     reflect.TypeOf(int8(0)),
			Value:    "10",
			Expected: reflect.ValueOf(int8(10)),
		},
		{
			Type:     reflect.TypeOf(int16(0)),
			Value:    "10",
			Expected: reflect.ValueOf(int16(10)),
		},
		{
			Type:     reflect.TypeOf(int32(0)),
			Value:    "10",
			Expected: reflect.ValueOf(int32(10)),
		},
		{
			Type:     reflect.TypeOf(int64(0)),
			Value:    "10",
			Expected: reflect.ValueOf(int64(10)),
		},
		{
			Type:     reflect.TypeOf(uint(0)),
			Value:    "100",
			Expected: reflect.ValueOf(uint(100)),
		},
		{
			Type:     reflect.TypeOf(uint8(0)),
			Value:    "100",
			Expected: reflect.ValueOf(uint8(100)),
		},
		{
			Type:     reflect.TypeOf(uint16(0)),
			Value:    "100",
			Expected: reflect.ValueOf(uint16(100)),
		},
		{
			Type:     reflect.TypeOf(uint32(0)),
			Value:    "100",
			Expected: reflect.ValueOf(uint32(100)),
		},
		{
			Type:     reflect.TypeOf(uint64(0)),
			Value:    "100",
			Expected: reflect.ValueOf(uint64(100)),
		},
		{
			Type:     reflect.TypeOf(float32(0.0)),
			Value:    "100.2",
			Expected: reflect.ValueOf(float32(100.2)),
		},
		{
			Type:     reflect.TypeOf(float64(0.0)),
			Value:    "100.2",
			Expected: reflect.ValueOf(float64(100.2)),
		},
		{
			Type:     reflect.TypeOf(false),
			Value:    "true",
			Expected: reflect.ValueOf(true),
		},
		{
			Type:     reflect.TypeOf(time.Time{}),
			Value:    "2020-01-02T03:04:05Z",
			Expected: reflect.ValueOf(time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC)),
		},
	}

	for _, testCase := range testCases {
		v, err := internal.Value2Var(testCase.Type, internal.Str(testCase.Value))
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected.Interface(), v.Interface())
	}
}
