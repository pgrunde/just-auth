package fields

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	day := NewDate(2015, 3, 1)
	assert.Equal(t, "2015-03-01", day.String())

	output, err := json.Marshal(day)
	assert.Nil(t, err, "JSON marshaling of dates should not error")
	assert.Equal(t, []byte(`"2015-03-01"`), output)

	nextDay := NewDate(2015, 3, 2)
	assert.True(t, nextDay.Equal(day.AddDays(1)))

	parsed, err := ParseDate("2015-03-01")
	assert.Nil(t, err, "Parsing of properly formatted dates should not error")
	assert.True(t, parsed.Equal(day))
}
