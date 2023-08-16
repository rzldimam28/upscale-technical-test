package logic

func CountLetterAndNumber(word string) (int, int) {

	var letterCount int
	var numberCount int

	for _, char := range word {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			letterCount++
		} else if char >= '0' && char <= '9' {
			numberCount++
		}
	}

	return letterCount, numberCount

}