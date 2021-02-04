package numbermanager_test

import (
	"testing"

	"github.com/maxxxlounge/interviews/SouthAfricanNumber/numbermanager"
	"github.com/stretchr/testify/assert"
)

func TestTrimNotNumbersDigit(t *testing.T) {
	tests := map[string]struct {
		input    string
		errMsg   string
		expected string
		hasError bool
	}{
		"EmptySpaces": {
			input:    " 1 ",
			expected: "1",
		},
		"StrangerChars": {
			input:    "AD12S!\"%&L120I ",
			expected: "12120",
		},
		"NothingToChange": {
			input:    "1234567891234",
			expected: "1234567891234",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out := numbermanager.TrimNotNumbersDigit(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}
}

func TestHasNotNumberDigits(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected bool
	}{
		"HasOnlyNumber": {
			input:    "12312312312",
			expected: false,
		},
		"HasNumberAndSpaces": {
			input:    "123 1231 2312",
			expected: true,
		},
		"HasOnlyNotNumbers": {
			input:    "asdasdsd asdasd",
			expected: true,
		},
		"MoreThanParseIntCheck": {
			input:    "12345678923456789123456789123459",
			expected: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out := numbermanager.HasNotNumberDigits(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}
}
