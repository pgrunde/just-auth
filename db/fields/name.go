package fields

import (
	"fmt"
	"strings"

	sql "github.com/aodin/aspect"
)

// MaxNameLength is the maximum number of characters for a name
const MaxNameLength = 64

type Name struct {
	Name string `db:"name" json:"name"`
}

// GetName returns the name
func (name Name) GetName() string {
	return name.Name
}

// SetName sets the name
func (name *Name) SetName(n string) {
	name.Name = n
}

// ErrorIfExists returns true if the table's column "name" already
// has a row that matches this value case in a case-insensitive manner
func (name Name) ErrorIfExists(conn sql.Connection, table *sql.TableElem) error {
	var duplicate string
	stmt := sql.Select(table.C["name"]).Where(
		sql.Lower(table.C["name"]).Equals(strings.ToLower(name.Name)),
	)
	if conn.MustQueryOne(stmt, &duplicate) {
		return fmt.Errorf(
			`A duplicate name "%s" already exists for "%s"`,
			duplicate, table.Name,
		)
	}
	return nil
}

func (name Name) ErrorIfExistsExcludingUUID(conn sql.Connection, table *sql.TableElem, uuid string) error {
	var duplicate string
	stmt := sql.Select(table.C["name"]).Where(
		sql.Lower(table.C["name"]).Equals(strings.ToLower(name.Name)),
		table.C["uuid"].DoesNotEqual(uuid),
	)
	if conn.MustQueryOne(stmt, &duplicate) {
		return fmt.Errorf(
			`A duplicate name "%s" already exists for "%s"`,
			duplicate, table.Name,
		)
	}
	return nil
}

func (name Name) Modify(table *sql.TableElem) error {
	return sql.Column("name", sql.String{NotNull: true}).Modify(table)
}

func (name Name) ErrorIfBlank() error {
	if name.Name == "" {
		return fmt.Errorf("Names cannot be blank")
	}
	return nil
}

func (name Name) ErrorIfTooLong() error {
	if len(name.Name) > MaxNameLength {
		return fmt.Errorf(
			"Names cannot be longer than %d characters", MaxNameLength,
		)
	}
	return nil
}

func (name Name) Error() error {
	if err := name.ErrorIfBlank(); err != nil {
		return err
	}
	if err := name.ErrorIfTooLong(); err != nil {
		return err
	}
	return nil
}

func NewName(name string) Name {
	return Name{Name: strings.TrimSpace(name)}
}
