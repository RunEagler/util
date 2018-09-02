package util

import (
	"time"
	"reflect"
)

func Value(x interface{}) reflect.Value {
	return reflect.ValueOf(x)
}

func Type(x interface{}) reflect.Type {
	return reflect.ValueOf(x).Type()
}

func PtrElem(x reflect.Value) reflect.Value {

	if x.Type().Kind() == reflect.Ptr {
		return x.Elem()
	}
	return x
}

func PtrInt(x int) *int {
	return &x
}

func PtrFloat(x float64) *float64 {
	return &x
}

func PtrString(x string) *string {
	return &x
}

func PtrTime(x time.Time) *time.Time {
	return &x
}

func PtrBool(x bool) *bool {
	return &x
}

func IsString(x reflect.Value) bool {

	return x.Kind() == reflect.String
}

func IsInt(x reflect.Value) bool {
	return x.Kind() == reflect.Int
}
