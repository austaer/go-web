package gotest

import (
	"errors"
)

func Division(divisor float64, dividend float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("Divisor cannot be zero.")
	}

	return divisor / dividend, nil
}

func Multiplication(multipler float64, multiplicand float64) float64 {
	return multipler * multiplicand
}
