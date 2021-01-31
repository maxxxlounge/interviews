package NumberManager_test

import (
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimNotNumbersDigit(t *testing.T) {
	assert.Equal(t, NumberManager.TrimNotNumbersDigit(" 1 "), "1")
	assert.Equal(t, NumberManager.TrimNotNumbersDigit(" AD12S!\"%&L120I "), "12120")
	assert.Equal(t, NumberManager.TrimNotNumbersDigit("1487751228"), "1487751228")
	assert.Equal(t, NumberManager.TrimNotNumbersDigit("_DELETED_1487751228"), "1487751228")
}

func TestIsRightFormat(t *testing.T) {
	assert.True(t,NumberManager.IsRightFormat("27831234567"))
	assert.False(t,NumberManager.IsRightFormat("27836826107_DELETED_1488996550"))
	assert.False(t,NumberManager.IsRightFormat("278368261071488996550"))
	assert.False(t,NumberManager.IsRightFormat("27841234567"))
	assert.False(t,NumberManager.IsRightFormat(""))
	assert.False(t,NumberManager.IsRightFormat("2784"))
}