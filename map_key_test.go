package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestMapStringKeys(t *testing.T) {

	testCases := []struct {
		values    interface{}
		expectVal []string
		expectErr error
	}{
		{
			values:    &map[string]int{"apple": 3, "banana": 5, "orange": 6},
			expectVal: []string{"apple", "banana", "orange"},
		},
		{
			values:    map[string]int{"apple": 3, "banana": 5, "orange": 6},
			expectVal: []string{"apple", "banana", "orange"},
		},

		{
			values:    map[string]bool{"apple": true, "banana": false, "orange": true},
			expectVal: []string{"apple", "banana", "orange"},
		},
		{
			values:    map[interface{}]interface{}{"apple": "basket", 5: nil, time.Now(): 0},
			expectVal: []string{},
		},
		{
			values:    map[int]int{1: 3, 2: 5, 3: 6},
			expectVal: []string{},
		},
		{
			values:    map[string]int{},
			expectVal: []string{},
		},
		{
			values:    10,
			expectVal: []string{},
		},
		{
			values:    true,
			expectVal: []string{},
		},
		{
			values:    nil,
			expectVal: []string{},
		},
	}
	for _, testCase := range testCases {
		actualVal := MapStringKeys(Value(testCase.values))
		assert.Equal(t, testCase.expectVal, actualVal, "expectVal should be actualErr")
	}
}

func TestMapIntKeys(t *testing.T) {

	testCases := []struct {
		values    interface{}
		expectVal []int
		expectErr error
	}{
		{
			values:    map[int]int{1: 3, 2: 5, 3: 6},
			expectVal: []int{1, 2, 3},
		},
		{
			values:    &map[int]int{1: 3, 2: 5, 3: 6},
			expectVal: []int{1, 2, 3},
		},
		{
			values:    map[int]string{0: "apple", 10: "banana", 8: "orange"},
			expectVal: []int{0, 8, 10},
		},
		{
			values:    &map[string]bool{"apple": true, "banana": false, "orange": true},
			expectVal: []int{},
		},
		{
			values:    map[interface{}]interface{}{"apple": "basket", 5: nil, time.Now(): 0},
			expectVal: []int{},
		},
		{
			values:    &map[string]int{},
			expectVal: []int{},
		},
		{
			values:    10,
			expectVal: []int{},
		},
		{
			values:    true,
			expectVal: []int{},
		},
		{
			values:    nil,
			expectVal: []int{},
		},
	}
	for _, testCase := range testCases {
		actualVal := MapIntKeys(Value(testCase.values))
		assert.Equal(t, testCase.expectVal, actualVal, "expectVal should be actualErr")
	}
}

func TestMapTimeKeys(t *testing.T) {

	testCases := []struct {
		values    interface{}
		expectVal []time.Time
		expectErr error
	}{
		{
			values:    map[time.Time]int{Date(2018, 5, 5): 0, Date(2000, 1, 1): 5, Datetime(2010, 9, 20, 18, 30): 10},
			expectVal: []time.Time{Date(2000, 1, 1), Datetime(2010, 9, 20, 18, 30), Date(2018, 5, 5)},
		},
		{
			values:    &map[time.Time]string{Date(2018, 5, 5): "error", Date(2000, 1, 1): "warning", Datetime(2010, 9, 20, 18, 30): "info"},
			expectVal: []time.Time{Date(2000, 1, 1), Datetime(2010, 9, 20, 18, 30), Date(2018, 5, 5)},
		},
		{
			values:    map[string]bool{"apple": true, "banana": false, "orange": true},
			expectVal: []time.Time{},
		},
		{
			values:    &map[interface{}]interface{}{"apple": "basket", 5: nil},
			expectVal: []time.Time{},
		},
		{
			values:    map[string]time.Time{},
			expectVal: []time.Time{},
		},
		{
			values:    10,
			expectVal: []time.Time{},
		},
		{
			values:    true,
			expectVal: []time.Time{},
		},
		{
			values:    nil,
			expectVal: []time.Time{},
		},
	}
	for _, testCase := range testCases {
		actualVal := MapTimeKeys(Value(testCase.values))
		assert.Equal(t, testCase.expectVal, actualVal, "expectVal should be actualErr")
	}
}

func TestMapStringIntSort(t *testing.T) {

	testCases := []struct {
		values    map[string]int
		expectVal []string
	}{
		{
			values:    map[string]int{"english": 30, "mathematics": 60, "japanese": 45, "physical": 20},
			expectVal: []string{"physical", "english", "japanese", "mathematics"},
		},
		{
			values:    map[string]int{},
			expectVal: []string{},
		},
		{
			values:    nil,
			expectVal: []string{},
		},
	}
	for _, testCase := range testCases {
		actualVal := MapStringIntSort(testCase.values)
		assert.Equal(t, testCase.expectVal, actualVal, "expectVal should be actualErr")
	}
}

func TestMapStringTimeSort(t *testing.T) {

	testCases := []struct {
		values    map[string]time.Time
		expectVal []string
	}{
		{
			values:    map[string]time.Time{"Smith": Date(1983, 4, 29), "John": Datetime(2011, 7, 15, 20, 5), "Ken": Date(1962, 10, 5), "Poal": Date(1974, 8, 20)},
			expectVal: []string{"Ken", "Poal", "Smith", "John"},
		},
		{
			values:    map[string]time.Time{},
			expectVal: []string{},
		},
		{
			values:    nil,
			expectVal: []string{},
		},
	}
	for _, testCase := range testCases {
		actualVal := MapStringTimeSort(testCase.values)
		assert.Equal(t, testCase.expectVal, actualVal, "expectVal should be actualErr")
	}
}
