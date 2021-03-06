package numbermanager_test

import (
	"fmt"
	"testing"

	"github.com/maxxxlounge/interviews/SouthAfricanNumber/numbermanager"
	"github.com/stretchr/testify/assert"
)

func TestFindCriticalError(t *testing.T) {
	errLessThanPrefix := fmt.Sprintf(numbermanager.ErrMsgLessThanPrefix, numbermanager.RightPrefix)
	errLessThanCore := fmt.Sprintf(numbermanager.ErrMsgLessThanCore, numbermanager.CoreLen)

	tests := map[string]struct {
		input    string
		errMsg   string
		hasError bool
	}{
		"NoDigit": {
			input:    "",
			errMsg:   errLessThanPrefix,
			hasError: true,
		},
		"LessThanPrefixDigit": {
			input:    "12",
			errMsg:   errLessThanPrefix,
			hasError: true,
		},
		"EqualThanPrefix": {
			input:    "1234",
			errMsg:   errLessThanCore,
			hasError: true,
		},
		"MoreThanPrefixLessThanCoreDigit": {
			input:    "12345",
			errMsg:   errLessThanCore,
			hasError: true,
		},
		"MoreEqualThanCoreDigit": {
			input:    "1234567",
			errMsg:   "",
			hasError: false,
		},
		"NoDigit_NotNumberMixed": {
			input:    "   ",
			errMsg:   errLessThanPrefix,
			hasError: true,
		},
		"LessThanPrefixDigit_NotNumberMixed": {
			input:    "12 _DELETED_",
			errMsg:   errLessThanPrefix,
			hasError: true,
		},
		"EqualThanPrefix_NotNumberMixed": {
			input:    "1234_ DELETED_",
			errMsg:   errLessThanCore,
			hasError: true,
		},
		"MoreThanPrefixLessThanCoreDigit_NotNumberMixed": {
			input:    " 12345_ASD!!",
			errMsg:   errLessThanCore,
			hasError: true,
		},
		"MoreEqualThanCoreDigit_NotNumberMixed": {
			input:    "  1234567  ASD!",
			errMsg:   "",
			hasError: false,
		},
		"MoreThanWantedDigit": {
			input:    "135449128449",
			errMsg:   "",
			hasError: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := numbermanager.FindCriticalError(tc.input)
			if tc.hasError {
				assert.EqualError(t, err, tc.errMsg)
			}
		})
	}
}
