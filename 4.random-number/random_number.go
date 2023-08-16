package random_number

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintRandomNumber(remaining map[int]bool, length int) {

	if length <= 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10) + 1

	if remaining[randomNumber] {
		PrintRandomNumber(remaining, length)
	} else {
		remaining[randomNumber] = true
		fmt.Printf("%d ", randomNumber)
		PrintRandomNumber(remaining, length-1)
	}

}