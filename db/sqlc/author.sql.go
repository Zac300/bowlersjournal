// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: author.sql

package db

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name, website, agent_id)
VALUES ($1, $2, $3)
RETURNING id, name, website, agent_id
`

type CreateAuthorParams struct {
	Name    string         `json:"name"`
	Website sql.NullString `json:"website"`
	AgentID int64          `json:"agent_id"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Website, arg.AgentID)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :one
DELETE FROM authors
WHERE id = $1
RETURNING id, name, website, agent_id
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, deleteAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return i, err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, website, agent_id FROM authors
WHERE id = $1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, website, agent_id FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
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

const listAuthorsByAgentID = `-- name: ListAuthorsByAgentID :many
SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, agents
WHERE agents.id = authors.agent_id AND authors.agent_id = $1
`

func (q *Queries) ListAuthorsByAgentID(ctx context.Context, agentID int64) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsByAgentID, agentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
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

const listAuthorsByBookID = `-- name: ListAuthorsByBookID :many
SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, book_authors
WHERE authors.id = book_authors.author_id AND book_authors.book_id = $1
`

func (q *Queries) ListAuthorsByBookID(ctx context.Context, bookID int64) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsByBookID, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
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

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors
SET name = $2, website = $3, agent_id = $4
WHERE id = $1
RETURNING id, name, website, agent_id
`

type UpdateAuthorParams struct {
	ID      int64          `json:"id"`
	Name    string         `json:"name"`
	Website sql.NullString `json:"website"`
	AgentID int64          `json:"agent_id"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, updateAuthor,
		arg.ID,
		arg.Name,
		arg.Website,
		arg.AgentID,
	)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return i, err
}