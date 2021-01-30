package NumberManager_test

import (
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimNotNumbersDigit(t *testing.T) {
	assert.Equal(t,NumberManager.TrimNotNumbersDigit(" 1 "),"1")
	assert.Equal(t,NumberManager.TrimNotNumbersDigit(" AD12S!\"%&L120I "),"12120")
	assert.Equal(t,NumberManager.TrimNotNumbersDigit("1487751228"),"1487751228")
	assert.Equal(t,NumberManager.TrimNotNumbersDigit("_DELETED_1487751228"),"1487751228")
}


