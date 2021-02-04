package numbermanager

import (
	"github.com/pkg/errors"
)

// FindCriticalError return critical error, that are numbers that is impossible to fix, 'cause some information missing.
// less valid int char than core number.
func FindCriticalError(input string) error {
	processedNumber := TrimNotNumbersDigit(input)
	processedNumberLen := len(processedNumber)

	if processedNumberLen < prefixLen {
		return errors.Errorf(ErrMsgLessThanPrefix, RightPrefix)
	}

	if processedNumberLen < CoreLen {
		return errors.Errorf(ErrMsgLessThanCore, CoreLen)
	}

	return nil
}
