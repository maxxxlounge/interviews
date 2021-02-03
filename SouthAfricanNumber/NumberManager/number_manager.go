package NumberManager

import (
	"log"
	"regexp"
)

type Row struct {
	Original string
	Changed  string
	Errors   []string `json:"Errors,omitempty"`
	Type     NumberType
}

func New(fullNumber string) *Row {
	r := Row{
		Original: fullNumber,
	}
	if IsRightFormat(fullNumber) {
		r.Type = ValidFirstAttempt
		return &r
	}
	err := FindCriticalError(fullNumber)
	if err != nil {
		r.Type = InvalidCritical
		r.Errors = append(r.Errors, err.Error())
		return &r
	}
	r.Type = InvalidButFixable
	r.Errors, r.Changed = FindFixableError(fullNumber)
	return &r
}

// HasNotNumberDigits verify if not number digits is present
// used despite of integer conversion check because the max was 4294967295
// If get and error during regex compile, it exit with fatal
func HasNotNumberDigits(number string) bool {
	hasNumberRegex, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return hasNumberRegex.MatchString(number)
}

// TrimNotNumbersDigit remove all not numeric digits
// If get and error during regex compile, it exit with fatal
func TrimNotNumbersDigit(input string) string {
	onlyNumberRegex, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	out := onlyNumberRegex.ReplaceAllString(input, "")
	if out != input {
		return out
	}
	return input
}
