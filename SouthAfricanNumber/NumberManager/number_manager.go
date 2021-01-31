package NumberManager

import (
	"github.com/pkg/errors"
	"log"
	"regexp"
	"strconv"
)

var CoreLen int = 7
var prefixLen int = 4

const RightPrefix string = "2783"

type NumberType string

const ValidFirstAttempt NumberType = "ValidFirstAttempt"
const InvalidCritical NumberType = "InvalidCritical"
const InvalidButFixable NumberType = "InvalidButFixable"

type Row struct {
	original string
	changed  string
	Errors   []error
	Type     NumberType
}

func (r *Row) GetChangedNumber() string {
	return r.changed
}


func (r *Row) GetOriginalNumber() string {
	return r.original
}

func New(fullNumber string) (*Row, error) {
	r := Row{
		original: fullNumber,
	}
	if IsRightFormat(fullNumber) {
		r.Type = ValidFirstAttempt
		return &r, nil
	}
	err := FindCriticalError(fullNumber)
	if err != nil {
		r.Type = InvalidCritical
		r.Errors = append(r.Errors, err)
		return &r, nil
	}
	r.Type = InvalidButFixable
	r.Errors,r.changed = FindFixableError(fullNumber)
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

func FindFixableError(number string) ([]error,string) {
	var changed string
	var errs []error
	var err error

	_, err = strconv.Atoi(number)
	if err != nil {
		err = errors.Errorf("error parsing '%s' to 0-9 digit, remove non number digits", number)
		errs = append(errs, err)
	}
	changed = TrimNotNumbersDigit(number)

	// missing prefix
	if len(changed)==CoreLen {
		changed = RightPrefix + changed
		err = errors.New("missing prefix, added one")
		errs = append(errs,err)
		return errs, changed
	}
	if len(changed)>CoreLen && len(changed)< CoreLen+prefixLen {
		changed = RightPrefix + changed
		err = errors.New("some number are missing added prefix")
		errs = append(errs,err)
	}

	if changed[0:3] != RightPrefix {
		changed = RightPrefix + changed[4:]
		err = errors.New("wrong prefix, replace with the correct one")
		errs = append(errs,err)
	}
	if len(changed)>CoreLen+prefixLen{
		changed = changed[0:CoreLen+prefixLen]
		err = errors.New("digit number more the wanted format, the exceeding was cut ")
		errs = append(errs,err)
	}

	// ok, from here we have a possible valid number for suggest correction
	/*prefix,err := GetPrefix(number)
	if err != nil {
		//ok, just to be sure that using GetPrefix err is managed in other situations
		errs = append(errs,err)
	}else {
		if prefix != RightPrefix {
			err = errors.Errorf("wrong prefix, needed '%s' given '%s'", RightPrefix, prefix)
			errs = append(errs,err)
		}
	}*/
	return errs,changed
}

func (r *Row) Correct(input string) (out string, err error) {
	return input, nil
}

func TrimNotNumbersDigit(input string) string {
	onlyNumberRegex, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return onlyNumberRegex.ReplaceAllString(input, "")

}
