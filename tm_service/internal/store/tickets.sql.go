// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tickets.sql

package store

import (
	"context"

	"github.com/google/uuid"
)

const createTicket = `-- name: CreateTicket :one
INSERT INTO tickets (user_id, title, content, priority)
VALUES ($1, $2, $3, $4)
RETURNING ticket_id
`

type CreateTicketParams struct {
	UserID   uuid.UUID `json:"user_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Priority string    `json:"priority"`
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.createTicketStmt, createTicket,
		arg.UserID,
		arg.Title,
		arg.Content,
		arg.Priority,
	)
	var ticket_id uuid.UUID
	err := row.Scan(&ticket_id)
	return ticket_id, err
}

const getTicketById = `-- name: GetTicketById :one
SELECT ticket_id, user_id, title, content, priority, status
FROM tickets
WHERE ticket_id = $1
`

func (q *Queries) GetTicketById(ctx context.Context, ticketID uuid.UUID) (Tickets, error) {
	row := q.queryRow(ctx, q.getTicketByIdStmt, getTicketById, ticketID)
	var i Tickets
	err := row.Scan(
		&i.TicketID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.Priority,
		&i.Status,
	)
	return i, err
}

const updateTicketStatus = `-- name: UpdateTicketStatus :exec
UPDATE tickets
SET status = $2
WHERE ticket_id = $1
`

type UpdateTicketStatusParams struct {
	TicketID uuid.UUID `json:"ticket_id"`
	Status   string    `json:"status"`
}

func (q *Queries) UpdateTicketStatus(ctx context.Context, arg UpdateTicketStatusParams) error {
	_, err := q.exec(ctx, q.updateTicketStatusStmt, updateTicketStatus, arg.TicketID, arg.Status)
	return err
}
