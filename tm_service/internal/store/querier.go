// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package store

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateTicket(ctx context.Context, arg CreateTicketParams) (uuid.UUID, error)
	GetTicketById(ctx context.Context, ticketID uuid.UUID) (Tickets, error)
	UpdateTicketStatus(ctx context.Context, arg UpdateTicketStatusParams) error
}

var _ Querier = (*Queries)(nil)