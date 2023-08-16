package palindrome

import (
	"strings"
)

func IsPalindrome(word string) bool {
	var reversedWord string

	// lowercasing
	lowerCaseWord := strings.ToLower(word)

	// remove white space
	cleanWord := strings.ReplaceAll(lowerCaseWord, " ", "")

	for i := len(cleanWord) - 1; i >= 0; i-- {
		reversedWord += string(cleanWord[i])
	}

	return cleanWord == reversedWord
}