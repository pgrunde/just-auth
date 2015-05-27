package fields

import (
	"time"

	sql "github.com/aodin/aspect"
	pg "github.com/aodin/aspect/postgres"
)

// Timestamp records creation, update, and optional deletion timestamps.
type Timestamp struct {
	CreatedAt time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

// Active returns true if the object has not been deleted
func (ts Timestamp) IsActive() bool {
	return ts.DeletedAt == nil
}

// Updated returns true if the timestamp has been updated
func (ts Timestamp) Updated() bool {
	return ts.UpdatedAt != nil
}

// Age returns the duration since the timestamp was created.
func (ts Timestamp) Age() time.Duration {
	return ts.age(time.Now().UTC())
}

func (ts Timestamp) age(now time.Time) time.Duration {
	return now.Sub(ts.CreatedAt)
}

// LastActivity returns the time of the lastest activity on the timestamp -
// either when it was last deleted, updated, or created
func (ts Timestamp) LastActivity() time.Time {
	if ts.DeletedAt != nil {
		return *ts.DeletedAt
	}
	if ts.UpdatedAt != nil {
		return *ts.UpdatedAt
	}
	return ts.CreatedAt
}

func (ts Timestamp) Modify(table *sql.TableElem) error {
	// TODO Determine the column names from the struct's db tags
	columns := []sql.ColumnElem{
		sql.Column("created_at", sql.Timestamp{NotNull: true, Default: pg.Now}),
		sql.Column("updated_at", sql.Timestamp{}),
		sql.Column("deleted_at", sql.Timestamp{}),
	}
	for _, column := range columns {
		if err := column.Modify(table); err != nil {
			return err
		}
	}
	return nil
}

// Timestamps should only be created by the database. This constructor should
// only be used for testing.
func newTimestamp(now time.Time) Timestamp {
	return Timestamp{CreatedAt: now}
}
