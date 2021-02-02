package NumberManager_test

import (
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**

 */
func TestAddDigitsWithPrefix(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
		hasError bool
		errMsg   string
	}{
		"NeedAddFullPrefix": {
			hasError: true,
			input:    "1234567",
			expected: NumberManager.RightPrefix + "1234567",
			errMsg:   NumberManager.ErrorMissingPartialPrefix,
		},
		"NeedAddPartialPrefix": {
			hasError: true,
			input:    "211234567",
			expected: NumberManager.RightPrefix[0:2] + "211234567",
			errMsg:   NumberManager.ErrorMissingPartialPrefix,
		},
		"NotNeedAddPrefix": {
			hasError: false,
			input:    "27831234567",
			expected: "27831234567",
			errMsg:   NumberManager.ErrorMissingPartialPrefix,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := NumberManager.AddDigitsWithPrefix(tc.input)
			assert.Equal(t, err != nil, tc.hasError)
			if err != nil {
				assert.Equal(t, tc.errMsg, err.Error())
			}
			assert.Equal(t, tc.expected, out)
		})
	}
}

func TestReplacePrefix(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
		hasError bool
		errMsg   string
	}{
		"ReplacePrefix": {
			hasError: true,
			input:    "12341234567",
			expected: NumberManager.RightPrefix + "1234567",
			errMsg:   NumberManager.ErrorWrongPrefix,
		},
		"ErrorLessThanPrefixLength": {
			hasError: true,
			input:    "123",
			expected: NumberManager.RightPrefix + "123",
			errMsg:   NumberManager.ErrMsgLessThanCore,
		},
		"ErrorLessThanCoreLength": {
			hasError: true,
			input:    "123456",
			expected: NumberManager.RightPrefix + "123456",
			errMsg:   NumberManager.ErrMsgLessThanCore,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := NumberManager.ReplacePrefix(tc.input)
			if err != nil {
				assert.Equal(t, tc.errMsg, err.Error())
			}
			assert.Equal(t, tc.expected, out)
		})
	}
}

func TestCutAdditionalDigits(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
		hasError bool
		errMsg   string
	}{
		"CutNotNecessary": {
			hasError: false,
			input:    "1234567",
			expected: "1234567",
			errMsg:   "",
		},
		"NeedAddPartialPrefix": {
			hasError: true,
			input:    "123412345678",
			expected: "12341234567",
			errMsg:   NumberManager.ErrorCutExtraDigits,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := NumberManager.CutAdditionalDigits(tc.input)
			assert.Equal(t, err != nil, tc.hasError)
			if err != nil {
				assert.Equal(t, tc.errMsg, err.Error())
			}
			assert.Equal(t, tc.expected, out)
		})
	}
}
