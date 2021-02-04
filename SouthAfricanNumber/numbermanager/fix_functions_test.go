package numbermanager_test

import (
	"testing"

	"github.com/maxxxlounge/interviews/SouthAfricanNumber/numbermanager"
	"github.com/stretchr/testify/assert"
)

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
			expected: numbermanager.RightPrefix + "1234567",
			errMsg:   numbermanager.ErrorMissingPartialPrefix,
		},
		"NeedAddPartialPrefix": {
			hasError: true,
			input:    "211234567",
			expected: numbermanager.RightPrefix[0:2] + "211234567",
			errMsg:   numbermanager.ErrorMissingPartialPrefix,
		},
		"NotNeedAddPrefix": {
			hasError: false,
			input:    "27831234567",
			expected: "27831234567",
			errMsg:   numbermanager.ErrorMissingPartialPrefix,
		},
		"NotOnlyNmberInputOK": {
			hasError: false,
			input:    "_DELECTED_27831234567_DELETED",
			expected: "27831234567",
			errMsg:   "",
		},
		"NotOnlyNmberInputWithErr": {
			hasError: true,
			input:    "_DELECTED_7831234567_DELETED",
			expected: "27831234567",
			errMsg:   numbermanager.ErrorMissingPartialPrefix,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := numbermanager.AddDigitsWithPrefix(tc.input)
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
			expected: numbermanager.RightPrefix + "1234567",
			errMsg:   numbermanager.ErrorWrongPrefix,
		},
		"NotReplacePrefix": {
			hasError: false,
			input:    "27831234567",
			expected: numbermanager.RightPrefix + "1234567",
			errMsg:   "",
		},
		"ErrorLessThanPrefixLength": {
			hasError: true,
			input:    "123",
			expected: "",
			errMsg:   numbermanager.ErrMsgLessThanCore,
		},
		"ErrorLessThanCoreLength": {
			hasError: true,
			input:    "123456",
			expected: "",
			errMsg:   numbermanager.ErrMsgLessThanCore,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := numbermanager.ReplacePrefix(tc.input)
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
			errMsg:   numbermanager.ErrorCutExtraDigits,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := numbermanager.CutAdditionalDigits(tc.input)
			assert.Equal(t, err != nil, tc.hasError)
			if err != nil {
				assert.Equal(t, tc.errMsg, err.Error())
			}
			assert.Equal(t, tc.expected, out)
		})
	}
}
