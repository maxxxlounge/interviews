package NumberManager

import (
	"github.com/pkg/errors"
)

const ErrMsgLessThanPrefix string =  "digits are less than wanted prefix '%s'"
const ErrMsgLessThanCore string = "digits are less than 'core' digits format '%v digits'"

/**
Critical error are numbers that is impossible to fix, 'cause some information missing
-- less valid int char than core number
*/
func FindCriticalError(input string)(error){
	processedNumber := TrimNotNumbersDigit(input)
	if len(processedNumber) < prefixLen{
		err := errors.Errorf(ErrMsgLessThanPrefix, RightPrefix)
		return err
	}
	if len(processedNumber) < CoreLen {
		err := errors.Errorf(ErrMsgLessThanCore,CoreLen)
		return err
	}
	return nil
}
