package main

import (
	"fmt"

	reading_file "github.com/rzldimam28/upscale-technical-test/1.reading-file"
	palindrome "github.com/rzldimam28/upscale-technical-test/2.palindrome"
	random_number "github.com/rzldimam28/upscale-technical-test/4.random-number"
	convert_time "github.com/rzldimam28/upscale-technical-test/5.convert-time"
	logic "github.com/rzldimam28/upscale-technical-test/6.logic"
)

func main() {
	reading_file.ReadingFile("test.txt")

	fmt.Println("PALINDROME")
	res := palindrome.IsPalindrome("Kasur ini rusak")
	fmt.Println(res)

	fmt.Println("RANDOM NUMBER")
	remaining := make(map[int]bool)
	length := 10
	random_number.PrintRandomNumber(remaining, length)
	
	fmt.Println("CONVERT TIME")
	value := "06:30:00 PM"
	t, err := convert_time.Convert12To24(value)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("time", t)
	
	value2 := "13:00:00"
	t2, err := convert_time.Convert12To24(value2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("time", t2)

	fmt.Println("FIZZBUZZ")
	// logic.FizzBuzz()

	fmt.Println("FACTORIAL")
	fmt.Println(logic.Factorial(5))

	fmt.Println("MINMAX")
	arr := []int{4, 2, 7, 10, 3, 14, 1, 12}
	minMaxArr, _ := logic.MinMax(arr)
	fmt.Println(minMaxArr)

	fmt.Println("COUNT CHAR")
	string1 := "Hello123World"
	string2 := "G0LangRocks"
	string3 := "OpenAI2023"
	lc1, nc1 := logic.CountLetterAndNumber(string1)
	lc2, nc2 := logic.CountLetterAndNumber(string2)
	lc3, nc3 := logic.CountLetterAndNumber(string3)
	fmt.Println(lc1, nc1)
	fmt.Println(lc2, nc2)
	fmt.Println(lc3, nc3)

}