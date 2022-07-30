package gowaybarplug

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xorcare/pointer"
)

func TestString(t *testing.T) {
	tests := []struct {
		input  Status
		output string
	}{{
		input:  Status{},
		output: `{"text":""}`,
	}, {
		input: Status{
			Text:       "a",
			Tooltip:    "b",
			Class:      []string{"c1", "c2"},
			Percentage: pointer.Int(42),
			Alt:        "d",
		},
		output: `{"text":"a","tooltip":"b","class":["c1","c2"],"percentage":42,"alt":"d"}`,
	}}
	for _, test := range tests {
		assert.Equal(t, test.input.String(), test.output)
	}
}
