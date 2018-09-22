package util

import (
	"reflect"
)

//TranspositionValue :interchanging the rows and columns by reflect.Value
func TranspositionValue(before [][]reflect.Value) [][]reflect.Value {

	var after [][]reflect.Value

	if len(before) == 0 {
		return before
	}
	after = make([][]reflect.Value, len(before[0]))
	for i := 0; i < len(before[0]); i++ {
		after[i] = make([]reflect.Value, len(before))
		for j := 0; j < len(before); j++ {
			after[i][j] = before[j][i]
		}
	}

	return after

}

//TranspositionInt :interchanging the rows and columns by int
func TranspositionInt(before [][]int) [][]int {

	var after [][]int

	if len(before) == 0 {
		return before
	}
	after = make([][]int, len(before[0]))
	for i := 0; i < len(before[0]); i++ {
		after[i] = make([]int, len(before))
		for j := 0; j < len(before); j++ {
			after[i][j] = before[j][i]
		}
	}

	return after

}

//TranspositionFloat64 :interchanging the rows and columns by float64
func TranspositionFloat64(before [][]float64) [][]float64 {

	var after [][]float64

	if len(before) == 0 {
		return before
	}
	after = make([][]float64, len(before[0]))
	for i := 0; i < len(before[0]); i++ {
		after[i] = make([]float64, len(before))
		for j := 0; j < len(before); j++ {
			after[i][j] = before[j][i]
		}
	}

	return after

}

//TranspositionFloat64 :interchanging the rows and columns by float64
func TranspositionString(before [][]string) [][]string {

	var after [][]string

	if len(before) == 0 {
		return before
	}
	after = make([][]string, len(before[0]))
	for i := 0; i < len(before[0]); i++ {
		after[i] = make([]string, len(before))
		for j := 0; j < len(before); j++ {
			after[i][j] = before[j][i]
		}
	}

	return after

}
