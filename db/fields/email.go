package fields

import (
	"fmt"
	"strings"

	sql "github.com/aodin/aspect"
)

type Email struct {
	Email string `db:"email" json:"email"`
}

// GetEmail returns the email
func (email Email) GetEmail() string {
	return email.Email
}

// SetEmail sets the email
func (email *Email) SetEmail(n string) {
	email.Email = n
}

func (email Email) Modify(table *sql.TableElem) error {
	return sql.Column("email", sql.String{NotNull: true}).Modify(table)
}

func (email Email) Error() error {
	// TODO is the best validation we have?
	parts := strings.Split(email.Email, "@")
	if len(parts) != 2 {
		return fmt.Errorf("There must be one and only one '@'")
	}
	if parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("Emails must be of the form 'user@domain'")
	}
	return nil
}

func NewEmail(email string) Email {
	return Email{Email: strings.TrimSpace(email)}
}
