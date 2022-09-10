package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	db "github.com/zac300/bowlersjournal/db/sqlc"
)

type Resolver struct {
	Repository db.Repository
}

// // foo
func (r *agentResolver) Authors(ctx context.Context, obj *db.Agent) ([]db.Author, error) {
	return r.Repository.ListAuthorsByAgentID(ctx, obj.ID)
}

// // foo
func (r *authorResolver) Website(ctx context.Context, obj *db.Author) (*string, error) {
	panic("not implemented")
}

// // foo
func (r *authorResolver) Agent(ctx context.Context, obj *db.Author) (*db.Agent, error) {
	panic("not implemented")
}

// // foo
func (r *authorResolver) Books(ctx context.Context, obj *db.Author) ([]db.Book, error) {
	panic("not implemented")
}

// // foo
func (r *bookResolver) Authors(ctx context.Context, obj *db.Book) ([]db.Author, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) CreateAgent(ctx context.Context, data AgentInput) (*db.Agent, error) {
	agent, err := r.Repository.CreateAgent(ctx, db.CreateAgentParams{
		Name:  data.Name,
		Email: data.Email,
	})
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

// // foo
func (r *mutationResolver) UpdateAgent(ctx context.Context, id int64, data AgentInput) (*db.Agent, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) DeleteAgent(ctx context.Context, id int64) (*db.Agent, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) CreateAuthor(ctx context.Context, data AuthorInput) (*db.Author, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) UpdateAuthor(ctx context.Context, id int64, data AuthorInput) (*db.Author, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id int64) (*db.Author, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) CreateBook(ctx context.Context, data BookInput) (*db.Book, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) UpdateBook(ctx context.Context, id int64, data BookInput) (*db.Book, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) DeleteBook(ctx context.Context, id int64) (*db.Book, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Agent(ctx context.Context, id int64) (*db.Agent, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Agents(ctx context.Context) ([]db.Agent, error) {
	return r.Repository.ListAgents(ctx)
}

// // foo
func (r *queryResolver) Author(ctx context.Context, id int64) (*db.Author, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Authors(ctx context.Context) ([]db.Author, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Book(ctx context.Context, id int64) (*db.Book, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Books(ctx context.Context) ([]db.Book, error) {
	panic("not implemented")
}

// Agent returns AgentResolver implementation.
func (r *Resolver) Agent() AgentResolver { return &agentResolver{r} }

// Author returns AuthorResolver implementation.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Book returns BookResolver implementation.
func (r *Resolver) Book() BookResolver { return &bookResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type agentResolver struct{ *Resolver }
type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
