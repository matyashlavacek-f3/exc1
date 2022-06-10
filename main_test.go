package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	type tc struct {
		input          string
		expectedOutput string
		expectedError  error
	}
	for _, scenario := range []tc{
		{
			input:          "",
			expectedOutput: "0",
		},
		{
			input:         "f",
			expectedError: fmt.Errorf("expected a number"),
		},
		{
			input:          "42",
			expectedOutput: "42",
		},
		{
			input:          "1+2",
			expectedOutput: "3",
		},
		{
			input:          "7+2",
			expectedOutput: "9",
		},
		{
			input:          "70+20",
			expectedOutput: "90",
		},
		{
			input:          "50-5",
			expectedOutput: "45",
		},
		{
			input:          "2 + 3",
			expectedOutput: "5",
		},
		{
			input:          "2 + 3 - 1",
			expectedOutput: "4",
		},
		{
			input:          "2 + 3 - ",
			expectedOutput: "5",
		},
		{
			input:          "2 + 4 * 5",
			expectedOutput: "22",
		},
		{
			input:          "2 + 10 / 2",
			expectedOutput: "7",
		},
		{
			input:          "2 + 2 * 5 * 2",
			expectedOutput: "22",
		},
		{
			input:          "-10",
			expectedOutput: "-10",
		},
		{
			input:          "5-10",
			expectedOutput: "-5",
		},
		{
			input:          "5*-10",
			expectedOutput: "-50",
		},
		{
			input:         "10/0",
			expectedError: fmt.Errorf("cannot divide by zero"),
		},
		{
			input:          "3/2",
			expectedOutput: "1.5",
		},
		{
			input:          "1/3",
			expectedOutput: "0.3333333333333333",
		},
	} {
		t.Run(scenario.input, func(t *testing.T) {
			output, err := Evaluate(scenario.input)
			if scenario.expectedError == nil {
				assert.NoError(t, err, "expected the evaluation to complete successfully")
				assert.Equal(t, scenario.expectedOutput, output, fmt.Sprintf("got %s, wanted %s", output, scenario.expectedOutput))
			} else {
				assert.EqualError(t, err, scenario.expectedError.Error())
			}
		})
	}
}
