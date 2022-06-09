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
	}
	for _, scenario := range []tc{
		{
			input:          "",
			expectedOutput: "0",
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
	} {
		t.Run(scenario.input, func(t *testing.T) {
			output := Evaluate(scenario.input)
			assert.Equal(t, scenario.expectedOutput, output, fmt.Sprintf("got %s, wanted %s", output, scenario.expectedOutput))
		})
	}
}
