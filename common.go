package util

import (
	"time"
	"reflect"
)

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

func IsString(x interface{}) bool {

	return reflect.ValueOf(x).Kind() == reflect.String
}

func IsInt(x interface{}) bool {
	return reflect.ValueOf(x).Kind() == reflect.Int
}
