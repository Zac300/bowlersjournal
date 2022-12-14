-- name: GetBook :one
SELECT * FROM books
WHERE id = $1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title;

-- name: ListBooksByAuthorID :many
SELECT books.* FROM books, book_authors
WHERE books.id = book_authors.book_id AND book_authors.author_id = $1;

-- name: CreateBook :one
INSERT INTO books (title, description, cover)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET title = $2, description = $3, cover = $4
WHERE id = $1
RETURNING *;

-- name: DeleteBook :one
DELETE FROM books
WHERE id = $1
RETURNING *;