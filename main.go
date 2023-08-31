package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Validator(phoneNumber string) (string, bool) {
	var errorMessage string
	var error bool

	if len(phoneNumber) != 10 {
		errorMessage = "Invald Length, Should be 10 charcters"
		error = true
		return errorMessage, error
	}

	firstDigit, _ := strconv.Atoi(string(phoneNumber[0]))
	fourthDigit, _ := strconv.Atoi(string(phoneNumber[3]))

	if firstDigit == 0 || firstDigit == 1 {
		errorMessage = "Number cannot start with 0 or 1"
		error = true
		return errorMessage, error
	}

	if fourthDigit == 0 || fourthDigit == 1 {
		errorMessage = "4th Character cannot start with 0 or 1"
		error = true
		return errorMessage, error
	}

	_, err := strconv.Atoi(phoneNumber)

	if err != nil {
		errorMessage = "Invalid with letters"
		error = true
		return errorMessage, error
	}

	error = false
	return "", error
}

func Number(phoneNumber string) (string, error) {

	phoneNumber = strings.TrimPrefix(phoneNumber, "+1")

	phoneNumber = strings.TrimPrefix(phoneNumber, "1")

	invalidChars := []string{"(", ")", "-", ".", " "}

	for _, num := range phoneNumber {
		for _, char := range invalidChars {
			if string(num) == char {
				before, after, _ := strings.Cut(phoneNumber, char)
				phoneNumber = before + after
			}
		}
	}

	errorMessage, error := Validator(phoneNumber)

	if !error {
		return phoneNumber, nil
	} else {
		return phoneNumber, fmt.Errorf(errorMessage)

	}

}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	var formatedNumber string

	number, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	formatedNumber = "(" + number[:3] + ")" + " " + number[3:6] + "-" + number[6:]

	return formatedNumber, nil
}
