// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package querytest

import (
	"database/sql"
)

type Bar struct {
	ID int64
}

type Foo struct {
	ID  int64
	Bar sql.NullInt64
}