// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package store

import (
	"context"
)

const listDummy = `-- name: ListDummy :many
SELECT id, name
FROM dummy
LIMIT 10
`

func (q *Queries) ListDummy(ctx context.Context) ([]Dummy, error) {
	rows, err := q.query(ctx, q.listDummyStmt, listDummy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dummy{}
	for rows.Next() {
		var i Dummy
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}