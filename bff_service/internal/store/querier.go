// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package store

import (
	"context"
)

type Querier interface {
	ListDummy(ctx context.Context) ([]Dummy, error)
}

var _ Querier = (*Queries)(nil)