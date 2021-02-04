package numbermanager

import (
	"log"
	"regexp"
)

// IsRightFormat return if the input match the right format
// return bool
func IsRightFormat(number string) bool {
	if len(number) != prefixLen+CoreLen {
		return false
	}
	matchedRegex, err := regexp.MatchString(RightPrefix+"[0-9]{7}", number)
	if err != nil {
		log.Fatal(err)
	}
	return matchedRegex
}
