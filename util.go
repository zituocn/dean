package dean

import (
	"reflect"
	"strconv"
)

func GetStructName(i interface{}) string {
	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

// Str2Int64 string to int64
func Str2Int64(s string) int64 {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i64
}

func Str2Int(s string) int {
	return int(Str2Int64(s))
}

// ToInt interface to int
//	default -1
func ToInt(v interface{}) int {
	i, ok := v.(int)
	if ok {
		return i
	}
	return -1
}

// ToInt64 interface to int64
//	default -1
func ToInt64(v interface{}) int64 {
	i, ok := v.(int64)
	if ok {
		return i
	}
	return -1
}

// ToFloat32 interface to float32
//	default -1
func ToFloat32(v interface{}) float32 {
	i, ok := v.(float32)
	if ok {
		return i
	}
	return -1
}

// ToFloat64 interface to float64
//	default -1
func ToFloat64(v interface{}) float64 {
	i, ok := v.(float64)
	if ok {
		return i
	}
	return -1
}

func ToString(v interface{}) string {
	i, ok := v.(string)
	if ok {
		return i
	}
	return ""
}
