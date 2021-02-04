package numbermanager_test

import (
	"testing"

	"github.com/maxxxlounge/interviews/SouthAfricanNumber/numbermanager"
	"github.com/stretchr/testify/assert"
)

func TestFindFixableError(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedErrMsg []string
		changed        string
		hasError       bool
	}{
		"TrimNotNumberTest": {
			input:          "_DELECTED_27831234567_DELETED",
			changed:        "27831234567",
			expectedErrMsg: []string{numbermanager.ErrorNotNumericDigits},
		},
		"NoDigit": {
			input:          "1234567",
			changed:        numbermanager.RightPrefix + "1234567",
			expectedErrMsg: []string{numbermanager.ErrorMissingPartialPrefix},
			hasError:       true,
		},
		"MorethanIntManage": {
			input:   "123123123123123123123",
			changed: numbermanager.RightPrefix + "2312312",
			expectedErrMsg: []string{
				numbermanager.ErrorCutExtraDigits,
				numbermanager.ErrorWrongPrefix,
			},
			hasError: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ee, changed := numbermanager.FindFixableError(tc.input)
			assert.Equal(t, tc.changed, changed)
			assert.Equal(t, len(tc.expectedErrMsg), len(ee), "has same errors number difference")
			for _, expectedError := range tc.expectedErrMsg {
				assert.Contains(t, ee, expectedError)
			}
		})
	}
}
