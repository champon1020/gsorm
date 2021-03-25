package internal

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/champon1020/mgorm/errors"
)

// BytesParser is parser for value whose type is []byte.
type BytesParser struct{}

// NewBytesParser creates BytesParser instance.
func NewBytesParser() *BytesParser {
	return &BytesParser{}
}

// BytesParserOption is option of BytesParser.
type BytesParserOption struct {
	TimeLayout string
}

// Parse convers bytes to target type.
// Supported types are as follows.
//  - string
//  - int, int8, int16, int32, int64
//  - uint, uint8, uint16, uint32, uint64
//  - float32, float64
//  - bool
//  - time.Time
func (p *BytesParser) Parse(bytes []byte, target reflect.Type, opt ...BytesParserOption) (*reflect.Value, error) {
	switch target.Kind() {
	case reflect.String:
		return p.String(bytes)
	case reflect.Int:
		return p.Int(bytes)
	case reflect.Int8:
		return p.Int8(bytes)
	case reflect.Int16:
		return p.Int16(bytes)
	case reflect.Int32:
		return p.Int32(bytes)
	case reflect.Int64:
		return p.Int64(bytes)
	case reflect.Uint:
		return p.Uint(bytes)
	case reflect.Uint8:
		return p.Uint8(bytes)
	case reflect.Uint16:
		return p.Uint16(bytes)
	case reflect.Uint32:
		return p.Uint32(bytes)
	case reflect.Uint64:
		return p.Uint64(bytes)
	case reflect.Float32:
		return p.Float32(bytes)
	case reflect.Float64:
		return p.Float64(bytes)
	case reflect.Bool:
		return p.Bool(bytes)
	case reflect.Struct:
		if target == reflect.TypeOf(time.Time{}) {
			layout := time.RFC3339
			if len(opt) != 0 && opt[0].TimeLayout != "" {
				layout = opt[0].TimeLayout
			}
			return p.Time(bytes, layout)
		}
	}

	msg := fmt.Sprintf("%s is not supported for BytesParser", target.String())
	return nil, errors.New(msg, errors.InvalidTypeError)
}

// String converts bytes to reflect.Value via string.
func (p *BytesParser) String(bytes []byte) (*reflect.Value, error) {
	v := reflect.ValueOf(string(bytes))
	return &v, nil
}

// Int converts bytes to reflect.Value via int.
func (p *BytesParser) Int(bytes []byte) (*reflect.Value, error) {
	i64, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(int(i64))
	return &v, nil
}

// Int8 converts bytes to reflect.Value via int8.
func (p *BytesParser) Int8(bytes []byte) (*reflect.Value, error) {
	i64, err := strconv.ParseInt(string(bytes), 10, 8)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(int8(i64))
	return &v, nil
}

// Int16 converts bytes to reflect.Value via int16.
func (p *BytesParser) Int16(bytes []byte) (*reflect.Value, error) {
	i64, err := strconv.ParseInt(string(bytes), 10, 16)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(int16(i64))
	return &v, nil
}

// Int32 converts bytes to reflect.Value via int32.
func (p *BytesParser) Int32(bytes []byte) (*reflect.Value, error) {
	i64, err := strconv.ParseInt(string(bytes), 10, 32)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(int32(i64))
	return &v, nil
}

// Int64 converts bytes to reflect.Value via int64.
func (p *BytesParser) Int64(bytes []byte) (*reflect.Value, error) {
	i64, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(i64)
	return &v, nil
}

// Uint converts bytes to reflect.Value via uint.
func (p *BytesParser) Uint(bytes []byte, bitSize ...int) (*reflect.Value, error) {
	u64, err := strconv.ParseUint(string(bytes), 10, 64)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(uint(u64))
	return &v, nil
}

// Uint8 converts bytes to reflect.Value via uint8.
func (p *BytesParser) Uint8(bytes []byte) (*reflect.Value, error) {
	u64, err := strconv.ParseUint(string(bytes), 10, 8)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(uint8(u64))
	return &v, nil
}

// Uint16 converts bytes to reflect.Value via uint16.
func (p *BytesParser) Uint16(bytes []byte) (*reflect.Value, error) {
	u64, err := strconv.ParseUint(string(bytes), 10, 16)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(uint16(u64))
	return &v, nil
}

// Uint32 converts bytes to reflect.Value via uint32.
func (p *BytesParser) Uint32(bytes []byte) (*reflect.Value, error) {
	u64, err := strconv.ParseUint(string(bytes), 10, 32)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(uint32(u64))
	return &v, nil
}

// Uint64 converts bytes to reflect.Value via uint64.
func (p *BytesParser) Uint64(bytes []byte) (*reflect.Value, error) {
	u64, err := strconv.ParseUint(string(bytes), 10, 64)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(u64)
	return &v, nil
}

// Float32 converts bytes to reflect.Value via float32.
func (p *BytesParser) Float32(bytes []byte) (*reflect.Value, error) {
	f64, err := strconv.ParseFloat(string(bytes), 32)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(float32(f64))
	return &v, nil
}

// Float64 converts bytes to reflect.Value via float64.
func (p *BytesParser) Float64(bytes []byte) (*reflect.Value, error) {
	f64, err := strconv.ParseFloat(string(bytes), 64)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(f64)
	return &v, nil
}

// Bool converts bytes to reflect.Value via bool.
func (p *BytesParser) Bool(bytes []byte) (*reflect.Value, error) {
	b, err := strconv.ParseBool(string(bytes))
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(b)
	return &v, nil
}

// Time converts bytes to reflect.Value via time.Time.
func (p *BytesParser) Time(bytes []byte, layout string) (*reflect.Value, error) {
	t, err := time.Parse(layout, string(bytes))
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(t)
	return &v, nil
}
