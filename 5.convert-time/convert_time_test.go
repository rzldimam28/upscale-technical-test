package convert_time_test

import (
	"testing"

	convert_time "github.com/rzldimam28/upscale-technical-test/5.convert-time"
	"github.com/stretchr/testify/assert"
)

func TestConvert12To24(t *testing.T) {

	testCases := []struct{
		Name string
		Error error
		Input string
		Output string
	}{
		{
			Name: "Success Convert Time",
			Error: nil,
			Input: "06:30:00 PM",
			Output: "18:30:00",
		},
		{
			Name: "Failed Convert Time | Invalid Input",
			Error: convert_time.ErrNotValidFormat,
			Input: "13:00:00",
			Output: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			output, err := convert_time.Convert12To24(tc.Input)

			assert.Equal(t, tc.Error, err)
			assert.Equal(t, tc.Output, output)
		})
	}

}