package logic_test

import (
	"testing"

	logic "github.com/rzldimam28/upscale-technical-test/6.logic"
	"github.com/stretchr/testify/assert"
)

func TestMinMax(t *testing.T) {

	testCases := []struct{
		Name string
		Input []int
		Output []int
	}{
		{
			Name: "Test Cases 1",
			Input: []int{42, 18, 90, 27, 71},
			Output: []int{18, 90},
		},
		{
			Name: "Test Cases 2",
			Input: []int{-10, 5, 32, -45, 17, 0, 23, -12},
			Output: []int{-45, 32},
		},
		{
			Name: "Test Cases 3",
			Input: []int{9, 2, 7, 0, 3, 8, 6, 1, 5, 4},
			Output: []int{0, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			output, err := logic.MinMax(tc.Input)

			assert.Nil(t, err)
			assert.Equal(t, tc.Output, output)
		})
	}

}