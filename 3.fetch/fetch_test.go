package fetch_test

import (
	"testing"

	fetch "github.com/rzldimam28/upscale-technical-test/3.fetch"
	"github.com/stretchr/testify/assert"
)

func TestFetchFromApi(t *testing.T) {

	testCases := []struct{
		Name string
		URL string
		Err error
		ResponseLength int
	}{
		{
			Name: "Success Fetch API",
			URL: fetch.URL,
			Err: nil,
			ResponseLength: 100,
		},
		{
			Name: "Failed Fetch API | Wrong URL",
			URL: "wrong-url.com",
			Err: fetch.ErrNotFoundURL,
			ResponseLength: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			respBody, err := fetch.FetchFromAPI(tc.URL)
		
			assert.Equal(t, tc.Err, err)
			assert.Equal(t, tc.ResponseLength, len(respBody))
		})	
	}
}