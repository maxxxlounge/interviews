package NumberManager_test

import (
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/stretchr/testify/assert"
	"testing"
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
			expected: NumberManager.RightPrefix + "1234567",
			errMsg:   NumberManager.ErrorMissingPartialPrefix,
		},
		"NeedAddPartialPrefix": {
			hasError: true,
			input:    "211234567",
			expected: NumberManager.RightPrefix + "27211234567",
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

}

func TestCutAdditionalDigits(t *testing.T) {

}
