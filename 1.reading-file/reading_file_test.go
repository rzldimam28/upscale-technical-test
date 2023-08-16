package reading_file_test

import (
	"testing"

	reading_file "github.com/rzldimam28/upscale-technical-test/1.reading-file"
	"github.com/stretchr/testify/assert"
)

func TestReadingFile(t *testing.T) {
	testCases := []struct{
		Name string
		FileName string
		ErrorExpected error
	}{
		{
			Name: "Success Reading File",
			FileName: "test.txt",
			ErrorExpected: nil,
		},
		{
			Name: "Failed Reading File | File Does Not Exist",
			FileName: "not-exist-file.txt",
			ErrorExpected: reading_file.ErrNotExisted,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := reading_file.ReadingFile(tc.FileName)

			assert.Equal(t, err, tc.ErrorExpected)
		})
	}
}