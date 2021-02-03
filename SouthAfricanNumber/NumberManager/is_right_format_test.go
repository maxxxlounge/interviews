package NumberManager_test

import (
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsRightFormat(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected bool
	}{
		"OK": {
			input:    "27831234567",
			expected: true,
		},
		"NotNumberFormat": {
			input:    "asdasdasd27831234567",
			expected: false,
		},
		"NotNumberFormatNumberFirst": {
			input:    "27831234567asdasdasd",
			expected: false,
		},
		"MoreThanMaxNumber": {
			input:    "278312345678",
			expected: false,
		},
		"LessThanMaxNumber": {
			input:    "2783123456",
			expected: false,
		},
		"WrongPrefix": {
			input:    "2784123456",
			expected: false,
		},
		"withspaces": {
			input:    "27841234 6",
			expected: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NumberManager.IsRightFormat(tc.input))
		})
	}

}
