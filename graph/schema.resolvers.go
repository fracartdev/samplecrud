package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/fracartdev/samplecrud/graph/generated"
	"github.com/fracartdev/samplecrud/graph/model"
)

func (r *mutationResolver) AddBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	id, err := r.Books.Create(input.Title, input.Author)
	if err != nil {
		return nil, err
	}

	log.Printf("Libro salvato con id: %s", *id)
	return &model.Book{
		ID:     *id,
		Title:  input.Title,
		Author: input.Author,
	}, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id string, updatedBook model.BookInput) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string, deletedBook model.BookInput) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReadBook(ctx context.Context, id string, readBook model.BookInput) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
