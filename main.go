package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Evaluate takes in a string representing a mathematical expression
// and returns the result
func Evaluate(input string) (string, error) {
	if input == "" {
		return "0", nil
	}

	input = strings.ReplaceAll(input, " ", "")

	sign := 1.0
	acc := ""
	operands := make([]float64, 0)
	operators := make([]rune, 0)
	for i, c := range input {
		if unicode.IsDigit(c) {
			acc = acc + string(c)
		} else if isOperator(c) {
			if i == 0 || !unicode.IsDigit(rune(input[i-1])) {
				sign = -1.0
				continue
			}

			n, err := parse(acc)
			if err != nil {
				return "", err
			}
			acc = ""
			operands = append(operands, n*sign)
			operators = append(operators, c)
			sign = 1.0
		} else {
			return "", fmt.Errorf("expected a number")
		}
	}

	if acc != "" {
		n, err := parse(acc)
		if err != nil {
			return "", err
		}
		operands = append(operands, n*sign)
	}

	operands, operators, err := resolvePriorityOps(operands, operators)
	if err != nil {
		return "", err
	}

	result := eval(operands, operators)
	return fmt.Sprint(result), nil
}

func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func resolvePriorityOps(operands []float64, operators []rune) ([]float64, []rune, error) {
	resOperands, resOperators := make([]float64, 0), make([]rune, 0)
	resOperands = append(resOperands, operands[0])

	if len(operands) == 1 {
		return operands, operators, nil
	}

	for i, op := range operators {
		if i+1 == len(operands) {
			break
		}
		switch op {
		case '*':
			resOperands[len(resOperands)-1] = resOperands[len(resOperands)-1] * operands[i+1]
		case '/':
			if operands[i+1] == 0 {
				return resOperands, resOperators, fmt.Errorf("cannot divide by zero")
			}
			resOperands[len(resOperands)-1] = resOperands[len(resOperands)-1] / operands[i+1]
		default:
			resOperands = append(resOperands, operands[i+1])
			resOperators = append(resOperators, operators[i])
		}
	}

	if operators[len(operators)-1] == '+' || operators[len(operators)-1] == '-' {
		resOperands = append(resOperands, operands[len(operands)-1])
	}

	return resOperands, resOperators, nil
}

func eval(operands []float64, operators []rune) float64 {
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

func parse(acc string) (float64, error) {
	if acc == "" {
		return 0, fmt.Errorf("expected a number")
	}
	n, err := strconv.ParseFloat(acc, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}
	return n, nil
}
