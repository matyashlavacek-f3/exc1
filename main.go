package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// Evaluate takes in a string representing a mathematical expression
// and returns the result
func Evaluate(input string) (string, error) {
	if input == "" {
		return "0", nil
	}
	acc := ""
	operands := make([]int, 0)
	operators := make([]rune, 0)
	for _, c := range input {
		if unicode.IsDigit(c) {
			acc = acc + string(c)
		} else if !unicode.IsSpace(c) {
			n, err := parse(acc)
			if err != nil {
				return "", err
			}
			acc = ""
			operands = append(operands, n)
			operators = append(operators, c)
		}
	}

	if acc != "" {
		n, err := parse(acc)
		if err != nil {
			return "", err
		}
		operands = append(operands, n)
	}

	result := eval(operands, operators)
	return fmt.Sprint(result), nil
}

func eval(operands []int, operators []rune) int {
	result := operands[0]
	for i, operator := range operators {
		if i+1 == len(operands) {
			break
		}
		switch operator {
		case '+':
			result = result + operands[i+1]
		case '-':
			result = result - operands[i+1]
		}
	}
	return result
}

func parse(acc string) (int, error) {
	n, err := strconv.Atoi(acc)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}
	return n, nil
}
