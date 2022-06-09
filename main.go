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
	operators := make([]rune, 0)
	for _, c := range input {
		if unicode.IsDigit(c) {
			acc = acc + string(c)
		} else if !unicode.IsSpace(c) {
			n := parse(acc)
			acc = ""
			operands = append(operands, n)
			operators = append(operators, c)
		}
	}

	if acc != "" {
		n := parse(acc)
		operands = append(operands, n)
	}

	result := operands[0]
	for i, operator := range operators {
		switch operator {
		case '+':
			result = result + operands[i+1]
		case '-':
			result = result - operands[i+1]
		}
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
