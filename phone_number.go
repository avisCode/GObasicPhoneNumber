package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

// Number the function returns a clean number if the number is a valid number
//and err with the relevant msg if the number is not valid
func Number(number string) (string, error) {
	cleanNumber := regexp.MustCompile("[^0-9]").ReplaceAllString(number, "")

	//Check length valid in case the length is 11 digits
	if len(cleanNumber) == 11 {
		if cleanNumber[0] != '1' {
			return "", errors.New("invalid format: if the number is longer then 10 digits the first number can be only 1")
		}
		cleanNumber = cleanNumber[1:]
	}

	//Check if length is 10 digits
	if len(cleanNumber) != 10 {
		return "", errors.New("invalid format: the number most be 10 digits")
	}

	// Check Area Code
	if cleanNumber[0] == '0' || cleanNumber[0] == '1' {
		return "", errors.New("invalid format: the Area Code can't start with 0 or 1")
	}

	// Check Exchange code
	if cleanNumber[3] == '0' || cleanNumber[3] == '1' {
		return "", errors.New("invalid format: the Exchange code can't start with 0 or 1")
	}

	return cleanNumber, nil
}

// AreaCode returns the area code in 3 digits like 917 if the given number is a valid number
//and err with the relevant msg if the number is not valid
func AreaCode(number string) (string, error) {
	if cleanNumber, err := Number(number); err != nil {
		return "", err
	} else {
		return cleanNumber[0:3], nil
	}
}

// Format returns the number in a format like (xxx) xxx-xxxx if given number is a valid number
//and err with the relevant msg if the number is not valid
func Format(number string) (string, error) {
	if cleanNumber, err := Number(number); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("(%s) %s-%s", cleanNumber[:3], cleanNumber[3:6], cleanNumber[6:]), nil
	}
}
