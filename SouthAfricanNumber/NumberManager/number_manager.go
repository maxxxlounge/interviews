package NumberManager

import (
	"github.com/pkg/errors"
	"log"
	"regexp"
)

const CoreLen int = 7
const prefixLen int = 4
const RightPrefix string = "2783"

type NumberType string

const ValidFirstAttempt NumberType = "ValidFirstAttempt"
const InvalidCritical NumberType = "InvalidCritical"
const InvalidButFixable NumberType = "InvalidButFixable"

type Row struct {
	Original string
	Changed  string
	Errors   []string `json:"string"`
	Type     NumberType
}

func New(fullNumber string) (*Row, error) {
	r := Row{
		Original: fullNumber,
	}
	if IsRightFormat(fullNumber) {
		r.Type = ValidFirstAttempt
		return &r, nil
	}
	err := FindCriticalError(fullNumber)
	if err != nil {
		r.Type = InvalidCritical
		r.Errors = append(r.Errors, err.Error())
		return &r, nil
	}
	r.Type = InvalidButFixable
	r.Errors,r.Changed = FindFixableError(fullNumber)
	return &r, nil
}

func IsRightFormat(number string) bool {
	if (len(number)!=prefixLen+CoreLen){
		return false
	}
	matchedRegex, err := regexp.MatchString(RightPrefix+"[0-9]{7}", number)
	if err != nil {
		err = errors.Wrap(err, "fatal occurred during regexp matching")
		log.Fatal(err)
	}
	return matchedRegex
}

func TrimNotNumbersDigit(input string) string {
	onlyNumberRegex, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return onlyNumberRegex.ReplaceAllString(input, "")
}
