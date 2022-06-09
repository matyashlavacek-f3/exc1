package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// Evaluate takes in a string representing a mathematical expression
// and returns the result
func Evaluate(input string) string {
	acc := ""
	result := 0
	for _, c := range input {
		if unicode.IsDigit(c) {
			acc = acc + string(c)
		} else {
			n := parse(acc)
			acc = ""
			result = result + n
		}
	}

	if acc != "" {
		n := parse(acc)
		result = result + n
	}

	return fmt.Sprint(result)
}

func parse(acc string) int {
	n, err := strconv.Atoi(acc)
	if err != nil {
		panic(err)
	}
	return n
}
