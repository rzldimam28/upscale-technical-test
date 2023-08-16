package palindrome_test

import (
	"testing"

	palindrome "github.com/rzldimam28/upscale-technical-test/2.palindrome"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {

	testCases := []struct{
		Name string
		Word string
	}{
		{
			Name: "IsPalindrome(Kasur ini rusak)",
			Word: "Kasur ini rusak",
		},
		{
			Name: "IsPalindrome(SAIPPUAKIVIKAUPPIAS)",
			Word: "SAIPPUAKIVIKAUPPIAS",
		},
		{
			Name: "IsPalindrome(Anna)",
			Word: "Anna",
		},
		{
			Name: "IsPalindrome(Civic)",
			Word: "Civic",
		},
		{
			Name: "IsPalindrome(My gym)",
			Word: "My gym",
		},
		{
			Name: "IsPalindrome(No lemon, no melon)",
			Word: "No lemon, no melon",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			isPal := palindrome.IsPalindrome(tc.Word)
			assert.True(t, isPal)
		})
	}

}