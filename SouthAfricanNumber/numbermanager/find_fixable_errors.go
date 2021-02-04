package numbermanager

/**
FindFixableError process record and give errors and fix suggestion
number is the given string
*/
func FindFixableError(number string) ([]string, string) {
	var (
		changed string
		errs    []string
		err     error
	)

	changed = number

	if HasNotNumberDigits(number) {
		errs = append(errs, ErrorNotNumericDigits)
		changed = TrimNotNumbersDigit(number)
	}

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
