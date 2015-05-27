package fields

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"fmt"
	"io"
	"log"

	sql "github.com/aodin/aspect"
	pg "github.com/aodin/aspect/postgres"
)

// Copyright 2011 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// UUID code is a variant of code.google.com/p/go-uuid
// Added database driver and restriction to v4

var UUIDv4 = pg.UUID{NotNull: true}

type uuid []byte

// String returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// , or "" if uuid is invalid.
func (uuid uuid) String() string {
	if len(uuid) != 16 {
		return ""
	}
	b := []byte(uuid)
	return fmt.Sprintf(
		"%08x-%04x-%04x-%04x-%012x", b[:4], b[4:6], b[6:8], b[8:10], b[10:],
	)
}

func (uuid uuid) MarshalJSON() ([]byte, error) {
	if len(uuid) == 0 {
		return []byte(`""`), nil
	}
	return []byte(`"` + uuid.String() + `"`), nil
}

func (u *uuid) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == `""` {
		return nil
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("UUIDs must be a valid UUID version 4")
	}
	data = data[1 : len(data)-1]
	uu, err := ParseUUID(string(data))
	if err != nil {
		return err
	}
	*u = uu.UUID
	return nil
}

// Scan converts an SQL value into a UUID
func (uuid *uuid) Scan(value interface{}) error {
	uu, _ := ParseUUID(string(value.([]byte)))
	*uuid = uu.UUID
	return nil
}

// Value returns the UUID formatted for insert into SQL
func (uuid uuid) Value() (driver.Value, error) {
	return uuid.String(), nil
}

type UUID struct {
	UUID uuid `db:"uuid" json:"uuid"`
}

// GetUUID returns the UUID as a string
func (uuid UUID) GetUUID() string {
	return uuid.String()
}

// Equals returns true if the UUIDs are equal
func (uuid UUID) Equals(other UUID) bool {
	return bytes.Equal(uuid.UUID, other.UUID)
}

// Exists returns true if the UUID is a valid UUID
func (uuid UUID) Exists() bool {
	return len(uuid.UUID) == 16
}

func (uuid UUID) Modify(table *sql.TableElem) error {
	return sql.Column("uuid", UUIDv4).Modify(table)
}

func (uuid UUID) String() string {
	return uuid.UUID.String()
}

// http://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_.28random.29
func NewUUID() (uuid UUID) {
	uuid.UUID = make([]byte, 16)
	randomBits([]byte(uuid.UUID))
	uuid.UUID[6] = (uuid.UUID[6] & 0x0f) | 0x40 // Version 4
	uuid.UUID[8] = (uuid.UUID[8] & 0x3f) | 0x80 // Variant is 10
	return
}

// TODO require v4
func ParseUUID(s string) (UUID, error) {
	if len(s) != 36 {
		return UUID{}, fmt.Errorf("UUIDs must have a length of 36 characters")
	}
	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return UUID{}, fmt.Errorf(
			"UUIDs must have a dash in the 8, 13, 18, and 23 positions",
		)
	}
	uuid := UUID{
		UUID: make([]byte, 16),
	}
	// Conver the hex characters to bytes
	for i, x := range []int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
		if v, ok := xtob(s[x:]); !ok {
			return UUID{}, fmt.Errorf(
				"UUIDs must have a valid hex encoded byte starting at position %d", x)
		} else {
			uuid.UUID[i] = v
		}
	}
	return uuid, nil
}

// randomBits completely fills slice b with random data.
func randomBits(b []byte) {
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Panic(err.Error()) // rand should never fail
	}
}

// xvalues returns the value of a byte as a hexadecimal digit or 255.
var xvalues = []byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// xtob converts the the first two hex bytes of x into a byte.
func xtob(x string) (byte, bool) {
	b1 := xvalues[x[0]]
	b2 := xvalues[x[1]]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}
