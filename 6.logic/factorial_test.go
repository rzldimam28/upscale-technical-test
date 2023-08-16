package logic_test

import (
	"testing"

	logic "github.com/rzldimam28/upscale-technical-test/6.logic"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {

	testCases := []struct{
		Name string
		Input int
		Output int
	}{
		{
			Name: "Factorial 5 Should Be 120",
			Input: 5,
			Output: 120,
		},
		{
			Name: "Factorial 3 Should Be 6",
			Input: 3,
			Output: 6,
		},
		{
			Name: "Factorial 8 Should Be 40320",
			Input: 8,
			Output: 40320,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			output := logic.Factorial(tc.Input)

			assert.Equal(t, tc.Output, output)
		})
	}

}