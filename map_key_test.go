package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"reflect"
)

func TestHasKey(t *testing.T) {

	testCases := []struct {
		values    interface{}
		key       interface{}
		expectVal bool
		expectErr error
	}{
		{
			values:    map[string]int{"a": 0, "aa": 1},
			key:       "a",
			expectVal: true,
			expectErr: nil,
		},
		/*
		{
			values:    map[string]interface{}{"a": 0, "aa": 1},
			key:       "a",
			expectVal: true,
			expectErr: nil,
		},
		{
			values:    map[string]interface{}{"a": 0, "aa": 1},
			key:       "a",
			expectVal: true,
			expectErr: nil,
		},
		{
			values:    map[string]interface{}{"a": 0, "aa": 1},
			key:       "a",
			expectVal: true,
			expectErr: nil,
		},
		{
			values:    map[string]interface{}{"a": 0, "aa": 1},
			key:       "a",
			expectVal: true,
			expectErr: nil,
		},
		{
			values:    map[string]interface{}{"a": 0, "aa": 1},
			key:       "a",
			expectVal: true,
			expectErr: nil,
		},
		{
			values:    map[string]interface{}{"a": 0, "aa": 1},
			key:       "a",
			expectVal: true,
			expectErr: nil,
		},
		*/
	}

	fmt.Println(reflect.TypeOf(hash).)

	for _, testCase := range testCases {
		actualVal := HasKey(testCase.values, testCase.key)
		assert.Equal(t, testCase.expectVal, actualVal, "expectVal should be actualErr")
		//assert.Equal(t, testCase.expectErr, actualErr, "expectErr should be actualErr")
	}

}
