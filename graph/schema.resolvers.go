package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fracartdev/samplecrud/graph/generated"
	"github.com/fracartdev/samplecrud/graph/model"
)

func (r *mutationResolver) AddBook(ctx context.Context, title *string, author *string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	var links []*model.Book
	dummyLink := model.Book{
		Title:  "our dummy link",
		Author: &model.Author{Name: "Dido"},
	}
	links = append(links, &dummyLink)
	return links, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
