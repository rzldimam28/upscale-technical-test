package convert_time

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrNotValidFormat = errors.New("input invalid, please use HH:MM:SS AM/PM format")
)

func inputValidation(timeIn12 string) error {
	if len(timeIn12) != 11 {
		return ErrNotValidFormat
	}

	parts := strings.Split(timeIn12, ":")
	hour := parts[0]
	minute := parts[1]
	second := parts[2][0:2]
	ampm := parts[2][3:5]

	hourInt, _ := strconv.Atoi(hour)
	minuteInt, _ := strconv.Atoi(minute)
	secondInt, _ := strconv.Atoi(second)

	if hourInt < 1 || hourInt > 12 {
		return ErrNotValidFormat
	}
	if minuteInt < 0 || minuteInt > 59 {
		return ErrNotValidFormat
	}
	if secondInt < 0 || secondInt > 59 {
		return ErrNotValidFormat
	}
	if ampm != "AM" && ampm != "PM" {
		return ErrNotValidFormat
	}

	return nil
}

func Convert12To24(timeIn12 string) (string, error) {

	err := inputValidation(timeIn12)
	if err != nil {
		return "", err
	}

	t, err := time.Parse("03:04:05 PM", timeIn12)
	if err != nil {
		return "", err
	}
	return t.Format("15:04:05"), nil
}

