package NumberManager_test

import (
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTrimNotNumbersDigitWithError(t *testing.T) {
	tests := map[string]struct {
		input    string
		errMsg   string
		expected string
		hasError bool
	}{
		"EmptySpaces": {
			input:    " 1 ",
			expected: "1",
			errMsg:   NumberManager.ErrorNotNumericDigits,
			hasError: true,
		},
		"StrangerChars": {
			input:    "AD12S!\"%&L120I ",
			expected: "12120",
			errMsg:   NumberManager.ErrorNotNumericDigits,
			hasError: true,
		},
		"NothingToChange": {
			input:    "1234567891234",
			expected: "1234567891234",
			errMsg:   "",
			hasError: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := NumberManager.TrimNotNumbersDigitWithError(tc.input)
			require.Equal(t, err != nil, tc.hasError)
			if tc.hasError {
				assert.Equal(t, tc.expected, out)
			}
		})
	}

}

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
			out := NumberManager.TrimNotNumbersDigit(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}

}

/*
func TestIsRightFormat(t *testing.T) {
	assert.True(t, NumberManager.IsRightFormat("27831234567"))
	assert.False(t, NumberManager.IsRightFormat("27836826107_DELETED_1488996550"))
	assert.False(t, NumberManager.IsRightFormat("278368261071488996550"))
	assert.False(t, NumberManager.IsRightFormat("27841234567"))
	assert.False(t, NumberManager.IsRightFormat(""))
	assert.False(t, NumberManager.IsRightFormat("2784"))
}
*/
