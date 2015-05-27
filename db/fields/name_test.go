package fields

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	hello := NewName("Hello")

	assert.Equal(t, "Hello", hello.GetName())
	assert.Nil(t, hello.Error(), "'Hello' should be a valid name")

	blank := NewName(" ")
	assert.NotNil(t, blank.Error(), "Blank names should error")

	tooLong := NewName(strings.Repeat("ab", MaxNameLength))
	assert.NotNil(t, tooLong.Error(), "Names longer than the max should error")
}
