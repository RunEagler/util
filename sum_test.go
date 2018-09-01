package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestSum(t *testing.T) {

	testCases := []struct {
		values    interface{}
		expectVal interface{}
		expectErr error
	}{
		{
			values:    []int{1, 2, 3, 4, 5, 6, 0},
			expectVal: 21,
			expectErr: nil,
		},
		{
			values:    []float64{1, 2.2, 3.0, 4.0, 0.0, 6.5},
			expectVal: 16.7,
			expectErr: nil,
		},
		{
			values:    []string{"abcde", "bcdef"},
			expectVal: nil,
			expectErr: fmt.Errorf("type of argument should be int or float64."),
		},
		{
			values:    nil,
			expectVal: nil,
			expectErr: fmt.Errorf("type of argument should be int or float64."),
		},
		{
			values:    []int{},
			expectVal: 0,
			expectErr: nil,
		},
		{
			values:    []float64{},
			expectVal: 0.0,
			expectErr: nil,
		},
	}

	for _, testCase := range testCases {
		actualVal, actualErr := Sum(testCase.values)
		assert.Equal(t, testCase.expectVal, actualVal, "expectVal should be actualErr")
		assert.Equal(t, testCase.expectErr, actualErr, "expectErr should be actualErr")
	}

}
