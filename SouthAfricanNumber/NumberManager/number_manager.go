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

type Row struct {
	original         string
	suggestedChanges string
	err              []error
}

func (r *Row) GetOriginalGivenNumber() string {
	return r.original
}

func New(fullNumber string) (*Row, error) {
	r := Row{
		original: fullNumber,
	}
	return &r, nil
}

/*
func GetPrefix(input string)(string,error){
	if (len(input) < prefixLen){
		return "",errors.Errorf("prefix short than %s",prefixLen)
	}
	return input[0:prefixLen-1],nil
}

func GetCore(input string)(string,error){
	if (len(input) < prefixLen+coreLen){
		return "",errors.Errorf("prefix short than %s",prefixLen)
	}
	return input[prefixLen:prefixLen+coreLen],nil
}*/

func FindFixableError(number string) ([]error) {
	var errs []error
	var err error
	_, err = strconv.Atoi(number)
	if err != nil {
		err = errors.Wrapf(err, "error parsing '%s' to 0-9 digit", number)
		errs = append(errs, err)
		return errs
	}
	matchedRegex, err := regexp.MatchString(RightPrefix + "[0-9]{7}", number)
	if err != nil {
		err = errors.Wrap(err, "error matching regex")
		errs = append(errs, err)
	}
	if !matchedRegex {
		err = errors.New("string doesn't match the correct format")
		errs = append(errs, err)
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
	return errs
}

func (r *Row) Correct(input string) (out string, err error) {
	return input, nil
}


func TrimNotNumbersDigit(input string)(string){
	onlyNumberRegex,err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return onlyNumberRegex.ReplaceAllString(input,"")

}