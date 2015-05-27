package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	email := NewEmail("  bill@murray.com ")
	assert.Equal(t, "bill@murray.com", email.GetEmail())

	errMail := NewEmail("hey@.@.com")
	assert.NotNil(t, errMail.Error())
}
