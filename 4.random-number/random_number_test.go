package random_number_test

import (
	"testing"

	random_number "github.com/rzldimam28/upscale-technical-test/4.random-number"
)

func TestPrintRandomNumber(t *testing.T) {
	remaining := make(map[int]bool)
	length := 10

	random_number.PrintRandomNumber(remaining, length)
}