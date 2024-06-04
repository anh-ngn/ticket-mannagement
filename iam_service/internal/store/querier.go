// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package store

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (GetUserByIDRow, error)
	GetUserByUsername(ctx context.Context, username string) (Users, error)
}

var _ Querier = (*Queries)(nil)
