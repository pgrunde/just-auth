package fields

import (
	"fmt"
	"strings"

	sql "github.com/aodin/aspect"
)

// MaxAboutLength specifies the maximum length for about
const MaxAboutLength = 1024

type About struct {
	About string `db:"about" json:"about"`
}

// GetAbout returns the about
func (about About) GetAbout() string {
	return about.About
}

// SetAbout sets the about
func (about *About) SetAbout(a string) {
	about.About = a
}

func (about About) Modify(table *sql.TableElem) error {
	return sql.Column("about", sql.String{}).Modify(table)
}

// Error returns an error if the about are invalid
func (about About) Error() error {
	if len(about.About) > MaxAboutLength {
		return fmt.Errorf(
			"About cannot be longer than %d characters", MaxAboutLength,
		)
	}
	return nil
}

// NewAbout creates a new About
func NewAbout(about string) About {
	return About{About: strings.TrimSpace(about)}
}
