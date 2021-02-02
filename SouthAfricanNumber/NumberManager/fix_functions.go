package NumberManager

import "errors"

const ErrorMissingPartialPrefix = "some number are missing added prefix"
const ErrorWrongPrefix = "wrong prefix, replace with the correct one"
const ErrorCutExtraDigits = "digit number more than wanted format, the exceeding was cut "
const ErrorNotNumericDigits = "found not numeric digits, removed"

func AddDigitsWithPrefix(number string) (string, error) {
	if len(number) >= CoreLen+prefixLen {
		return number, nil
	}
	missingNumberCount := (CoreLen + prefixLen) - len(number)
	number = RightPrefix[0:missingNumberCount] + number
	err := errors.New(ErrorMissingPartialPrefix)
	return number, err
}

func ReplacePrefix(number string) (string, error) {
	if len(number) < CoreLen {
		return RightPrefix + number, errors.New(ErrMsgLessThanCore)
	}
	if number[0:3] == RightPrefix {
		return number, nil
	}
	number = RightPrefix + number[4:]
	err := errors.New(ErrorWrongPrefix)
	return number, err
}

func CutAdditionalDigits(number string) (string, error) {
	if len(number) <= CoreLen+prefixLen {
		return number, nil
	}
	number = number[0 : CoreLen+prefixLen]
	err := errors.New(ErrorCutExtraDigits)
	return number, err
}
