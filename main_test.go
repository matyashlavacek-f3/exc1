package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	type tc struct {
		input          string
		expectedOutput string
	}
	for _, scenario := range []tc{
		{
			input:          "1+2",
			expectedOutput: "3",
		},
		{
			input:          "7+2",
			expectedOutput: "9",
		},
	} {
		t.Run(scenario.input, func(t *testing.T) {
			output := Evaluate(scenario.input)
			assert.Equal(t, scenario.expectedOutput, output, fmt.Sprintf("expected the result of the addition to be %s", scenario.expectedOutput))
		})
	}
}
