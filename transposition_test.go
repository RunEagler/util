package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func intToValueSlice(before []int) []reflect.Value {

	var after []reflect.Value

	for _, val := range before {
		after = append(after, Value(val))
	}
	return after
}
func stringToValueSlice(before []string) []reflect.Value {

	var after []reflect.Value

	for _, val := range before {
		after = append(after, Value(val))
	}
	return after
}

func TestTranspositionValue(t *testing.T) {

	testCases := []struct {
		values    [][]reflect.Value
		expectVal [][]reflect.Value
	}{
		{
			values: [][]reflect.Value{
				intToValueSlice([]int{1, 2, 3, 4, 5}),
				intToValueSlice([]int{1, 2, 3, 4, 5}),
				intToValueSlice([]int{1, 2, 3, 4, 5}),
				intToValueSlice([]int{1, 2, 3, 4, 5}),
				intToValueSlice([]int{1, 2, 3, 4, 5}),
			},
			expectVal: [][]reflect.Value{
				intToValueSlice([]int{1, 1, 1, 1, 1}),
				intToValueSlice([]int{1, 2, 2, 2, 2}),
				intToValueSlice([]int{3, 3, 3, 3, 3}),
				intToValueSlice([]int{4, 4, 4, 4, 4}),
				intToValueSlice([]int{5, 5, 5, 5, 5}),
			},
		},
		{
			values: [][]reflect.Value{
				intToValueSlice([]int{1, 2, 3, 4, 5}),
				intToValueSlice([]int{1, 2, 3, 4, 5}),
			},
			expectVal: [][]reflect.Value{
				intToValueSlice([]int{1, 1}),
				intToValueSlice([]int{2, 2}),
				intToValueSlice([]int{3, 3}),
				intToValueSlice([]int{4, 4}),
				intToValueSlice([]int{5, 5}),
			},
		},
		{
			values: [][]reflect.Value{
				stringToValueSlice([]string{"a", "b", "c", "d"}),
			},
			expectVal: [][]reflect.Value{
				stringToValueSlice([]string{"a"}),
				stringToValueSlice([]string{"b"}),
				stringToValueSlice([]string{"c"}),
				stringToValueSlice([]string{"d"}),
			},
		},
		{
			values: [][]reflect.Value{
				intToValueSlice([]int{1, 2, 3, 4, 5}),
			},
			expectVal: [][]reflect.Value{
				intToValueSlice([]int{1}),
				intToValueSlice([]int{2}),
				intToValueSlice([]int{3}),
				intToValueSlice([]int{4}),
				intToValueSlice([]int{5}),
			},
		},
		{
			values: [][]reflect.Value{
				intToValueSlice([]int{1}),
			},
			expectVal: [][]reflect.Value{
				intToValueSlice([]int{1}),
			},
		},
		{
			values: [][]reflect.Value{
				intToValueSlice([]int{}),
			},
			expectVal: [][]reflect.Value{},
		},
		{
			values:    [][]reflect.Value{},
			expectVal: [][]reflect.Value{},
		},
	}

	for _, testCase := range testCases {
		actualVal := TranspositionValue(testCase.values)
		equalValue2Slice(t, testCase.expectVal, actualVal)
	}

}

func TestTranspositionInt(t *testing.T) {

	testCases := []struct {
		values    [][]int
		expectVal [][]int
	}{
		{
			values: [][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			},
			expectVal: [][]int{
				{1, 1, 1, 1, 1},
				{2, 2, 2, 2, 2},
				{3, 3, 3, 3, 3},
				{4, 4, 4, 4, 4},
				{5, 5, 5, 5, 5},
			},
		},
		{
			values: [][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			},
			expectVal: [][]int{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{5, 5},
			},
		},
		{
			values: [][]int{
				{1, 2, 3, 4, 5},
			},
			expectVal: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
		},
		{
			values: [][]int{
				{1},
			},
			expectVal: [][]int{
				{1},
			},
		},
		{
			values: [][]int{
				{},
			},
			expectVal: [][]int{},
		},
		{
			values:    [][]int{},
			expectVal: [][]int{},
		},
	}

	for _, testCase := range testCases {
		actualVal := TranspositionInt(testCase.values)
		assert.Equal(t, testCase.expectVal, actualVal)
	}

}

func TestTranspositionFloat64(t *testing.T) {

	testCases := []struct {
		values    [][]float64
		expectVal [][]float64
	}{
		{
			values: [][]float64{
				{1.1, 2.2, 3, 4, 5},
				{1.1, 2.2, 3, 4, 5},
				{1.1, 2.2, 3, 4, 5},
				{1.1, 2.2, 3, 4, 5},
				{1.1, 2.2, 3, 4, 5},
			},
			expectVal: [][]float64{
				{1.1, 1.1, 1.1, 1.1, 1.1},
				{2.2, 2.2, 2.2, 2.2, 2.2},
				{3, 3, 3, 3, 3},
				{4, 4, 4, 4, 4},
				{5, 5, 5, 5, 5},
			},
		},
		{
			values: [][]float64{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			},
			expectVal: [][]float64{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{5, 5},
			},
		},
		{
			values: [][]float64{
				{1, 2, 3, 4, 5},
			},
			expectVal: [][]float64{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
		},
		{
			values: [][]float64{
				{1},
			},
			expectVal: [][]float64{
				{1},
			},
		},
		{
			values: [][]float64{
				{},
			},
			expectVal: [][]float64{},
		},
		{
			values:    [][]float64{},
			expectVal: [][]float64{},
		},
	}

	for _, testCase := range testCases {
		actualVal := TranspositionFloat64(testCase.values)
		assert.Equal(t, testCase.expectVal, actualVal)
	}

}

func TestTranspositionString(t *testing.T) {

	testCases := []struct {
		values    [][]string
		expectVal [][]string
	}{
		{
			values: [][]string{
				{"1", "2", "3", "4", "5"},
				{"1", "2", "3", "4", "5"},
				{"1", "2", "3", "4", "5"},
				{"1", "2", "3", "4", "5"},
				{"1", "2", "3", "4", "5"},
			},
			expectVal: [][]string{
				{"1", "1", "1", "1", "1"},
				{"2", "2", "2", "2", "2"},
				{"3", "3", "3", "3", "3"},
				{"4", "4", "4", "4", "4"},
				{"5", "5", "5", "5", "5"},
			},
		},
		{
			values: [][]string{
				{"1", "2", "3", "4", "5"},
				{"1", "2", "3", "4", "5"},
			},
			expectVal: [][]string{
				{"1", "1"},
				{"2", "2"},
				{"3", "3"},
				{"4", "4"},
				{"5", "5"},
			},
		},
		{
			values: [][]string{
				{"1", "2", "3", "4", "5"},
			},
			expectVal: [][]string{
				{"1"},
				{"2"},
				{"3"},
				{"4"},
				{"5"},
			},
		},
		{
			values: [][]string{
				{"1"},
			},
			expectVal: [][]string{
				{"1"},
			},
		},
		{
			values: [][]string{
				{},
			},
			expectVal: [][]string{},
		},
		{
			values:    [][]string{},
			expectVal: [][]string{},
		},
	}

	for _, testCase := range testCases {
		actualVal := TranspositionString(testCase.values)
		assert.Equal(t, testCase.expectVal, actualVal)
	}

}

func equalValue2Slice(t *testing.T, expect, actual [][]reflect.Value) {

	assert.Equal(t, len(expect), len(actual), "len(columns)")
	for i, expectRow := range expect {
		actualRow := actual[i]
		assert.Equal(t, len(expectRow), len(actualRow), "len(rows)")
		for j, expectCell := range expectRow {
			actualCell := actualRow[j]
			assert.Equal(t, expectCell.String(), actualCell.String(), "Transpositon")
		}
	}

}
