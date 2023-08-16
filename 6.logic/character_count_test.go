package logic_test

import (
	"testing"

	logic "github.com/rzldimam28/upscale-technical-test/6.logic"
	"github.com/stretchr/testify/assert"
)

func TestCountLetterAndNumber(t *testing.T) {

	testCases := []struct{
		Name string
		Input string
		LetterCount int
		NumberCount int
	}{
		{
			Name: "Test Cases 1",
			Input: "Hello123World",
			LetterCount: 10,
			NumberCount: 3,
		},
		{
			Name: "Test Cases 2",
			Input: "G0LangRocks",
			LetterCount: 10,
			NumberCount: 1,
		},
		{
			Name: "Test Cases 3",
			Input: "OpenAI2023",
			LetterCount: 6,
			NumberCount: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			letterCount, numberCount := logic.CountLetterAndNumber(tc.Input)

			assert.Equal(t, tc.LetterCount, letterCount)
			assert.Equal(t, tc.NumberCount, numberCount)
		})
	}

}