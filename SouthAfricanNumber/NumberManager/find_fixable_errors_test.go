package NumberManager_test

import (
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFixableError(t *testing.T) {
	ee,fixed := NumberManager.FindFixableError("123123123123123123123")
	assert.Equal(t,"error parsing '123123123123123123123' to 0-9 digit, remove non number digits",)
}
