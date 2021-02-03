package NumberManager

import "errors"


// AddDigitsWithPrefix add missing prefix digits before number if needed
// If number digits is more than prefix+core len exit
// If prefix digit added, it returns a error
// return changed number and error if changes occurred
func AddDigitsWithPrefix(number string) (string, error) {
	if HasNotNumberDigits(number){
		number = TrimNotNumbersDigit(number)
	}
	if len(number) >= CoreLen+prefixLen {
		return number, nil
	}
	missingNumberCount := (CoreLen + prefixLen) - len(number)
	number = RightPrefix[0:missingNumberCount] + number
	err := errors.New(ErrorMissingPartialPrefix)
	return number, err
}

// ReplacePrefix identify and replace wrong prefix if needed
// return changed number and error if changes occurred
func ReplacePrefix(number string) (string, error) {
	if HasNotNumberDigits(number){
		number = TrimNotNumbersDigit(number)
	}
	if len(number) < CoreLen {
		return "", errors.New(ErrMsgLessThanCore)
	}
	if number[0:prefixLen] == RightPrefix {
		return number, nil
	}
	number = RightPrefix + number[prefixLen:]
	err := errors.New(ErrorWrongPrefix)
	return number, err
}

// CutAdditionalDigits remove the extra numeric digits at the end of input
// return changed number and error if changes occurred
func CutAdditionalDigits(number string) (string, error) {
	if HasNotNumberDigits(number){
		number = TrimNotNumbersDigit(number)
	}
	if len(number) <= CoreLen+prefixLen {
		return number, nil
	}
	number = number[0 : CoreLen+prefixLen]
	err := errors.New(ErrorCutExtraDigits)
	return number, err
}
