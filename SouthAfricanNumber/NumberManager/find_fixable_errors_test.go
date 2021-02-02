package NumberManager_test

import (
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/stretchr/testify/assert"
	"testing"
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
			expectedErrMsg: []string{NumberManager.ErrorNotNumericDigits},
		},
		"NoDigit": {
			input:          "1234567",
			changed:        NumberManager.RightPrefix + "1234567",
			expectedErrMsg: []string{NumberManager.ErrorMissingPartialPrefix},
			hasError:       true,
		},
		"MorethanIntManage": {
			input:          "123123123123123123123",
			changed:        NumberManager.RightPrefix + "2312312",
			expectedErrMsg: []string{NumberManager.ErrorCutExtraDigits},
			hasError:       true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ee, changed := NumberManager.FindFixableError(tc.input)
			assert.Equal(t, len(tc.expectedErrMsg), len(ee), "has same errors number difference")
			for _, expectedError := range tc.expectedErrMsg {
				assert.Contains(t, ee, expectedError)
			}
			assert.Equal(t, tc.changed, changed)
		})
	}
}
