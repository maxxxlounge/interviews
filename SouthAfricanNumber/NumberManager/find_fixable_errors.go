package NumberManager

import (
	"github.com/pkg/errors"
	"strconv"
)

func FindFixableError(number string) ([]string, string) {
	var changed string
	var errs []string
	var err error

	_, err = strconv.Atoi(number)
	if err != nil {
		err = errors.Errorf("error parsing '%s' to 0-9 digit, remove non number digits", number)
		errs = append(errs, err.Error())
	}
	changed = TrimNotNumbersDigit(number)
	changed, err = AddDigitsWithPrefix(changed)
	if err != nil {
		errs = append(errs, err.Error())
	}
	changed, err = ReplacePrefix(changed)
	if err != nil {
		errs = append(errs, err.Error())
	}
	changed, err = CutAdditionalDigits(changed)
	if err != nil {
		errs = append(errs, err.Error())
	}
	return errs, changed
}
