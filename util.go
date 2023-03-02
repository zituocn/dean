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
