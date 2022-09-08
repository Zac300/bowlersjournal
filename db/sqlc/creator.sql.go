// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: creator.sql

package db

import (
	"context"
)

const setBookAuthor = `-- name: SetBookAuthor :exec
INSERT INTO book_authors (book_id, author_id)
VALUES ($1, $2)
`

type SetBookAuthorParams struct {
	BookID   int64 `json:"book_id"`
	AuthorID int64 `json:"author_id"`
}

func (q *Queries) SetBookAuthor(ctx context.Context, arg SetBookAuthorParams) error {
	_, err := q.db.ExecContext(ctx, setBookAuthor, arg.BookID, arg.AuthorID)
	return err
}

const unsetBookAuthors = `-- name: UnsetBookAuthors :exec
DELETE FROM book_authors
WHERE book_id = $1
`

func (q *Queries) UnsetBookAuthors(ctx context.Context, bookID int64) error {
	_, err := q.db.ExecContext(ctx, unsetBookAuthors, bookID)
	return err
}
