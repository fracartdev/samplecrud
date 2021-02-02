package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/fracartdev/samplecrud/graph/generated"
	"github.com/fracartdev/samplecrud/graph/model"
)

func (r *mutationResolver) AddBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	var book model.Book

	book.Title = input.Title
	book.Author = input.Author

	return &book, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	dummyBook := model.Book{
		Title:  "Reti Logiche",
		Author: "Dido",
	}
	books = append(books, &dummyBook)
	return books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
