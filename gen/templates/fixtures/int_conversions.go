package fixtures

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidNumber = errors.New("invalid Number")
)

func IsNumber(in int) bool {
	switch in {
	case 1:
		return true
	case 2:
		return true
	case 3:
		return true
	}

	return false
}

func ToNumber(in int) (Number, bool) {
	switch in {
	case 1:
		return NumberOne, true
	case 2:
		return NumberTwo, true
	case 3:
		return NumberThree, true
	}

	return Number(0), false
}

func ToNumberErr(in int) (Number, error) {
	if number, ok := ToNumber(in); ok {
		return number, nil
	}

	return Number(0), fmt.Errorf("casting `%v`: %w", in, ErrInvalidNumber)
}

func MustToNumber(in int) Number {
	number, err := ToNumberErr(in)
	if err != nil {
		panic(err)
	}

	return number
}
