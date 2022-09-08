// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
)

type Agent struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Author struct {
	ID      int64          `json:"id"`
	Name    string         `json:"name"`
	Website sql.NullString `json:"website"`
	AgentID int64          `json:"agent_id"`
}

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}

type BookAuthor struct {
	ID       int64 `json:"id"`
	BookID   int64 `json:"book_id"`
	AuthorID int64 `json:"author_id"`
}
