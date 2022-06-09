package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// Evaluate takes in a string representing a mathematical expression
// and returns the result
func Evaluate(input string) string {
	if input == "" {
		return "0"
	}
	acc := ""
	operands := make([]int, 0)
	var operator rune
	for _, c := range input {
		if unicode.IsDigit(c) {
			acc = acc + string(c)
		} else {
			n := parse(acc)
			acc = ""
			operands = append(operands, n)
			operator = c
		}
	}

	if acc != "" {
		n := parse(acc)
		operands = append(operands, n)
	}

	var result int
	switch operator {
	case '+':
		result = operands[0] + operands[1]
	case '-':
		result = operands[0] - operands[1]
	default:
		result = operands[0]
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
