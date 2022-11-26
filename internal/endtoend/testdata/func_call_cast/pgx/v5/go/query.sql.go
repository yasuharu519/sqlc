// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const demo = `-- name: Demo :one
SELECT uuid_generate_v5('7c4597a0-8cfa-4c19-8da0-b8474a36440d', $1)::uuid as col1
`

func (q *Queries) Demo(ctx context.Context, uuidGenerateV5 interface{}) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, demo, uuidGenerateV5)
	var col1 pgtype.UUID
	err := row.Scan(&col1)
	return col1, err
}