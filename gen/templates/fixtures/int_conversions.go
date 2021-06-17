package fixtures

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidNumber = errors.New("invalid Number")
)

func IsNumber(s int) bool {
	switch s {
	case 1:
		return true
	case 2:
		return true
	case 3:
		return true
	}

	return false
}

func ToNumber(s int) (Number, bool) {
	switch s {
	case 1:
		return NumberOne, true
	case 2:
		return NumberTwo, true
	case 3:
		return NumberThree, true
	}

	return Number(0), false
}

func ToNumberErr(s int) (Number, error) {
	if number, ok := ToNumber(s); ok {
		return number, nil
	}

	return Number(0), fmt.Errorf("casting `%v`: %w", s, ErrInvalidNumber)
}

func MustToNumber(s int) Number {
	number, err := ToNumberErr(s)
	if err != nil {
		panic(err)
	}

	return number
}
