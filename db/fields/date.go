package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"time"
)

// DateFormat is the standard date format for display and parsing
const DateFormat = "2006-01-02"

type Date struct{ time.Time }

func (date Date) format() string {
	return date.Time.Format(DateFormat)
}

// AddDays adds the given number of days to the date
func (date Date) AddDays(days int) Date {
	return Date{Time: date.Time.AddDate(0, 0, days)}
}

// String returns the Date as a string
func (date Date) String() string {
	return date.format()
}

// Equal returns true if the dates are equal
func (date Date) Equal(other Date) bool {
	return date.Time.Equal(other.Time)
}

// UnmarshalJSON converts a byte array into a Date
func (d *Date) UnmarshalJSON(text []byte) error {
	b := bytes.NewBuffer(text)
	dec := json.NewDecoder(b)
	var s string
	if err := dec.Decode(&s); err != nil {
		return err
	}
	value, err := time.Parse(DateFormat, s)
	if err != nil {
		return err
	}
	d.Time = value
	return nil
}

// MarshalJSON returns the JSON output of a Date
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.format() + `"`), nil
}

// Scan converts an SQL value into a Date
func (date *Date) Scan(value interface{}) error {
	date.Time = value.(time.Time)
	return nil
}

// Value returns the date formatted for insert into SQL
func (date Date) Value() (driver.Value, error) {
	return date.format(), nil
}

// Today converts the local time to a Date - as more campuses open this will
// need to be timezone aware
func Today() Date {
	return today(time.Now())
}

func today(now time.Time) Date {
	return NewDate(now.Year(), now.Month(), now.Day())
}

// NewDate creates a new date.
func NewDate(year int, month time.Month, day int) Date {
	// Remove all second and nano second information and mark as UTC
	return Date{
		Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

// ParseDate converts a date string to a Date, possibly returning an error
func ParseDate(value string) (Date, error) {
	return ParseDateWithFormat(DateFormat, value)
}

// ParseDateWithFormat calls ParseDate with a different date format
func ParseDateWithFormat(format, value string) (Date, error) {
	t, err := time.Parse(format, value)
	if err != nil {
		return Date{}, err
	}
	return Date{Time: t}, nil
}
