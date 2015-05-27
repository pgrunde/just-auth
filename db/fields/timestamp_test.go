package fields

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestamp(t *testing.T) {
	now := time.Date(2015, 3, 2, 0, 0, 0, 0, time.UTC)
	then := time.Date(2015, 3, 1, 0, 0, 0, 0, time.UTC)

	ts := newTimestamp(then)
	assert.True(t, ts.IsActive(), "Timestamp should be active")
	assert.False(t, ts.Updated(), "Timestamp should not be updated")
	assert.Equal(t,
		24*time.Hour,
		ts.age(now),
		"Timestamp should be one day old",
	)

	ts.DeletedAt = &now
	assert.False(t, ts.IsActive(), "Timestamp should be inactive")
}
