package util

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestCombination(t *testing.T) {

	testCases := []struct {
		n      int
		r      int
		expect [][]int
	}{
		{
			n: 4,
			r: 2,
			expect: [][]int{
				{1, 0}, {2, 0}, {2, 1}, {3, 0}, {3, 1}, {3, 2},
			},
		},
		{
			n: 5,
			r: 3,
			expect: [][]int{
				{2, 1, 0}, {3, 1, 0}, {3, 2, 0}, {3, 2, 1}, {4, 1, 0},
				{4, 2, 0}, {4, 2, 1}, {4, 3, 0}, {4, 3, 1}, {4, 3, 2},
			},
		},
		{
			n: 0,
			r: 0,
			expect: [][]int{
				{0},
			},
		},
		{
			n: 5,
			r: 0,
			expect: [][]int{
				{0},
			},
		},
		{
			n:      0,
			r:      5,
			expect: [][]int{},
		},
		{
			n:      1,
			r:      5,
			expect: [][]int{},
		},
		{
			n: 5,
			r: 1,
			expect: [][]int{
				{0}, {1}, {2}, {3}, {4},
			},
		},
	}

	for No, testCase := range testCases {
		var actual [][]int
		actual = [][]int{}
		combination(testCase.n, testCase.r, []int{}, &actual)
		assert.Equal(t, testCase.expect, actual, fmt.Sprintf("No = %d", No))
	}

}
