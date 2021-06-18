package fixtures

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrInvalidCity = errors.New("invalid City")
)

func IsCity(in string) bool {
	switch in {
	case "london":
		return true
	case "oakland":
		return true
	case "portland":
		return true
	case "seattle":
		return true
	case "San Francisco":
		return true
	case `"`:
		return true
	case "sekret":
		return true
	}

	return false
}

func ToCity(in string) (City, bool) {
	switch in {
	case "london":
		return CityLondon, true
	case "oakland":
		return CityOakland, true
	case "portland":
		return CityPortland, true
	case "seattle":
		return CitySeattle, true
	case "San Francisco":
		return CitySanFrancisco, true
	case `"`:
		return CityQuotes, true
	case "sekret":
		return citySekret, true
	}

	return City(""), false
}

func ToCityErr(in string) (City, error) {
	if city, ok := ToCity(in); ok {
		return city, nil
	}

	return City(""), fmt.Errorf("casting `%v`: %w", in, ErrInvalidCity)
}

func MustToCity(in string) City {
	city, err := ToCityErr(in)
	if err != nil {
		panic(err)
	}

	return city
}

type CityValues struct{}

func (_ *CityValues) Values() []string {
	return []string{
		"london",
		"oakland",
		"portland",
		"seattle",
		"San Francisco",
		`"`,
		"sekret",
	}
}

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

func ParseNumber(in string) (Number, error) {
	num, err := strconv.Atoi(in)
	if err != nil {
		return Number(0), fmt.Errorf("parsing `%v` to number: %w", in, err)
	}

	result, err := ToNumberErr(num)
	if err != nil {
		return Number(0), err
	}

	return result, nil
}

type NumberValues struct{}

func (_ *NumberValues) Values() []int {
	return []int{
		1,
		2,
		3,
	}
}
