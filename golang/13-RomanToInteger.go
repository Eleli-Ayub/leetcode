package main

import (
	"strings"
)

func RomanToInteger(s string) int {
	romanValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	x := len(s) - 1
	s = strings.ToUpper(s)
	totalValue := 0
	currentValue := 0
	prevValue := 0
	for x >= 0 {
		if x == len(s)-1 {
			prevValue = 0
		} else {
			prevValue, _ = romanValues[string(s[x+1])]
		}

		currentValue, _ = romanValues[string(s[x])]

		if currentValue < prevValue {
			totalValue -= currentValue
		} else {
			totalValue += currentValue
		}

		x--
	}
	return totalValue
}
