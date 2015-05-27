package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	uuid := NewUUID()

	// Generated UUIDs should be valid
	parsed, err := ParseUUID(uuid.String())
	assert.Nil(t, err)
	assert.True(t, uuid.Equals(parsed))
	assert.True(t, uuid.Exists())
}
