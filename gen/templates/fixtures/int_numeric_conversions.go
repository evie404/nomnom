package fixtures

import (
	"fmt"
	"strconv"
)

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
