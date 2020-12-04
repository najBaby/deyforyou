package handler

import (
	"fmt"
	"net/url"
	"strconv"
)

const form = "form: field' %s' is empty"

// Form is
type Form struct {
	value url.Values
}

// NewForm is
func NewForm(v url.Values) *Form {
	return &Form{
		value: v,
	}
}

// Int is
func (f *Form) Int(name string) (int, error) {
	i64, err := f.Int64(name)
	return int(i64), err
}

// Ints is
func (f *Form) Ints(name string) ([]int, error) {
	i64s, err := f.Int64s(name)
	result := make([]int, 0, len(i64s))
	fmt.Println(result)
	for _, i := range i64s {
		result = append(result, int(i))
	}
	return result, err
}

// Int8 is
func (f *Form) Int8(name string) (int8, error) {
	i64, err := f.Int64(name)
	return int8(i64), err
}

// Int8s is
func (f *Form) Int8s(name string) ([]int8, error) {
	i64s, err := f.Int64s(name)
	result := make([]int8, 0, len(i64s))
	for _, i := range i64s {
		result = append(result, int8(i))
	}
	return result, err
}

// Int16 is
func (f *Form) Int16(name string) (int16, error) {
	i64, err := f.Int64(name)
	return int16(i64), err
}

// Int16s is
func (f *Form) Int16s(name string) ([]int16, error) {
	i64s, err := f.Int64s(name)
	result := make([]int16, 0, len(i64s))
	for _, i := range i64s {
		result = append(result, int16(i))
	}
	return result, err
}

// Int32 is
func (f *Form) Int32(name string) (int32, error) {
	i64, err := f.Int64(name)
	return int32(i64), err
}

// Int32s is
func (f *Form) Int32s(name string) ([]int32, error) {
	i64s, err := f.Int64s(name)
	result := make([]int32, 0, len(i64s))
	for _, i := range i64s {
		result = append(result, int32(i))
	}
	return result, err
}

// Int64 is
func (f *Form) Int64(name string) (int64, error) {
	if v := f.value.Get(name); v != "" {
		return strconv.ParseInt(v, 10, 64)
	}
	return 0, fmt.Errorf(form, name)
}

// Int64s is
func (f *Form) Int64s(name string) ([]int64, error) {
	var result []int64
	if v := f.value[name]; v != nil {
		for _, v := range v {
			i, e := strconv.ParseInt(v, 10, 64)
			if e != nil {
				return result, e
			}
			result = append(result, i)
		}
		return result, nil
	}
	return result, fmt.Errorf(form, name)
}

// Uint is
func (f *Form) Uint(name string) (uint, error) {
	i64, err := f.Uint64(name)
	return uint(i64), err
}

// Uints is
func (f *Form) Uints(name string) ([]uint, error) {
	i64s, err := f.Uint64s(name)
	result := make([]uint, 0, len(i64s))
	for _, i := range i64s {
		result = append(result, uint(i))
	}
	return result, err
}

// Uint8 is
func (f *Form) Uint8(name string) (uint8, error) {
	i64, err := f.Uint64(name)
	return uint8(i64), err
}

// Uint8s is
func (f *Form) Uint8s(name string) ([]uint8, error) {
	i64s, err := f.Uint64s(name)
	result := make([]uint8, 0, len(i64s))
	for _, i := range i64s {
		result = append(result, uint8(i))
	}
	return result, err
}

// Uint16 is
func (f *Form) Uint16(name string) (uint16, error) {
	i64, err := f.Uint64(name)
	return uint16(i64), err
}

// Uint16s is
func (f *Form) Uint16s(name string) ([]uint16, error) {
	i64s, err := f.Uint64s(name)
	result := make([]uint16, 0, len(i64s))
	for _, i := range i64s {
		result = append(result, uint16(i))
	}
	return result, err
}

// Uint32 is
func (f *Form) Uint32(name string) (uint32, error) {
	i64, err := f.Uint64(name)
	return uint32(i64), err
}

// Uint32s is
func (f *Form) Uint32s(name string) ([]uint32, error) {
	i64s, err := f.Uint64s(name)
	result := make([]uint32, 0, len(i64s))
	for _, i := range i64s {
		result = append(result, uint32(i))
	}
	return result, err
}

// Uint64 is
func (f *Form) Uint64(name string) (uint64, error) {
	if v := f.value.Get(name); v != "" {
		return strconv.ParseUint(v, 10, 64)
	}
	return 0, fmt.Errorf(form, name)
}

// Uint64s is
func (f *Form) Uint64s(name string) ([]uint64, error) {
	var result []uint64
	if v := f.value[name]; v != nil {
		for _, v := range v {
			i, e := strconv.ParseUint(v, 10, 64)
			if e != nil {
				return result, e
			}
			result = append(result, i)
		}
		return result, nil
	}
	return result, fmt.Errorf(form, name)
}

// Float32 is
func (f *Form) Float32(name string) (float32, error) {
	i64, err := f.Float64(name)
	return float32(i64), err
}

// Float32s is
func (f *Form) Float32s(name string) ([]float32, error) {
	i64s, err := f.Float64s(name)
	result := make([]float32, 0, len(i64s))
	for _, i := range i64s {
		result = append(result, float32(i))
	}
	return result, err
}

// Float64 is
func (f *Form) Float64(name string) (float64, error) {
	if v := f.value.Get(name); v != "" {
		return strconv.ParseFloat(v, 64)
	}
	return 0, fmt.Errorf(form, name)
}

// Float64s is
func (f *Form) Float64s(name string) ([]float64, error) {
	var result []float64
	if v := f.value[name]; v != nil {
		for _, v := range v {
			i, e := strconv.ParseFloat(v, 64)
			if e != nil {
				return result, e
			}
			result = append(result, i)
		}
		return result, nil
	}
	return result, fmt.Errorf(form, name)
}

// Bool is
func (f *Form) Bool(name string) (bool, error) {
	v := f.value.Get(name)
	if v != "" {
		return strconv.ParseBool(v)
	}
	return false, fmt.Errorf(form, name)
}

// String is
func (f *Form) String(name string) (string, error) {
	v := f.value.Get(name)
	if v != "" {
		return v, nil
	}
	return v, fmt.Errorf(form, name)
}

// Strings is
func (f *Form) Strings(name string) ([]string, error) {
	v := f.value[name]
	if v != nil {
		return v, nil
	}
	return v, fmt.Errorf(form, name)
}
