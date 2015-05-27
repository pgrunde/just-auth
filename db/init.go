package db

import (
	sql "github.com/aodin/aspect"
	pg "github.com/aodin/aspect/postgres"
)

// Schema
var Users, Sessions *sql.TableElem

func init() {
	Users = sql.Table("users",
		sql.Column("id", pg.Serial{NotNull: true}),
		fields.Name{},
		fields.Email{},
		fields.Timestamp{},
		fields.About{},
		sql.Column("is_active", sql.Boolean{NotNull: true, Default: sql.False}),
		sql.Column("is_superuser", sql.Boolean{NotNull: true, Default: sql.False}),
		fields.Password{},
		sql.Column("token", sql.String{NotNull: true}),
		sql.Column("token_set_at", sql.Timestamp{NotNull: true, Default: pg.Now}),
		sql.PrimaryKey("id"),
	)
	Sessions = sql.Table("sessions",
		sql.Column("key", sql.String{NotNull: true}),
		sql.ForeignKey("user_id", Users.C["id"], sql.Integer{NotNull: true}).OnDelete(sql.Cascade),
		sql.Column("expires", sql.Timestamp{NotNull: true}),
	)
}
