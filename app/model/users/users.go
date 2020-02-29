package users

import "github.com/gogf/gf/container/gset"

const (
	// ColumnID - the id column of entity.
	ColumnID = "id"
	// ColumnName - the name column of entity.
	ColumnName = "name"
	// ColumnEmail - the email column of entity.
	ColumnEmail = "email"
	// ColumnPhone - the phone column of entity.
	ColumnPhone = "phone"
	// ColumnPassword - the password column of entity.
	ColumnPassword = "password"
	// ColumnRating - the rating column of entity.
	ColumnRating = "rating"
)

var (
	// AllQueryColumns - all columns used for query, which doesn't include sensitive columns (e.g. password).
	AllQueryColumns = gset.NewFrom([]string{ColumnID, ColumnName, ColumnEmail, ColumnPhone, ColumnRating})
)
