package fixtures

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidCity = errors.New("invalid City")
)

func IsCity(s string) bool {
	switch s {
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

func ToCity(s string) (City, bool) {
	switch s {
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

func ToCityErr(s string) (City, error) {
	if city, ok := ToCity(s); ok {
		return city, nil
	}

	return City(""), fmt.Errorf("casting `%v`: %w", s, ErrInvalidCity)
}

func MustToCity(s string) City {
	city, err := ToCityErr(s)
	if err != nil {
		panic(err)
	}

	return city
}
