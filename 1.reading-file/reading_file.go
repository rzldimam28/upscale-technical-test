package reading_file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
)

var (
	ErrNotExisted = errors.New("file does not exist")
)

func ReadingFile(fileName string) error {

	filePath := "D:/session/go/src/upscale-technical-test/1.reading-file/"

	pathToFile := path.Join(filePath, fileName)

	file, err := os.Open(pathToFile)
	if err != nil {
		return ErrNotExisted
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var scanToArray []int
	for scanner.Scan() {
		line := scanner.Text()
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		scanToArray = append(scanToArray, lineInt)
	}

	var sum int
	for _, value := range scanToArray {
		sum += value
	}

	fmt.Printf("Total angka pada file: %d\n", len(scanToArray))
	fmt.Printf("Jumlah semua angka: %d\n", sum)

	return nil
}