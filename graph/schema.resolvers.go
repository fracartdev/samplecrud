package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/fracartdev/samplecrud/books"
	"github.com/fracartdev/samplecrud/graph/generated"
	"github.com/fracartdev/samplecrud/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) AddBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	id, err := r.BooksLib.Create(input.Title, input.Author)
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
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("id non valido: %s", id)
	}

	err = r.BooksLib.Update(id, updatedBook.Title, updatedBook.Author)
	if err != nil {
		return nil, err
	}

	log.Printf("Libro con id: %s aggiornato", id)
	return &model.Book{
		ID:     id,
		Title:  updatedBook.Title,
		Author: updatedBook.Author,
	}, nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string, deletedBook model.BookInput) (*model.Book, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("id non valido: %s", id)
	}

	err = r.BooksLib.Delete(id)
	if err != nil {
		return nil, err
	}

	log.Printf("Libro con id: %s eliminato", id)
	return &model.Book{
		ID:     id,
		Title:  deletedBook.Title,
		Author: deletedBook.Author,
	}, nil
}

func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("id non valido: %s", id)
	}

	result, err := r.BooksLib.Read(id)
	if err != nil {
		return nil, err
	}

	log.Printf("Libro con id: %s trovato", id)
	return &model.Book{
		ID:     id,
		Title:  result.Title,
		Author: result.Author,
	}, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	var items []*model.Book
	var savedItems []books.BookItem
	savedItems, err := r.BooksLib.List()
	if err != nil {
		return nil, err
	}
	for i, savedItem := range savedItems {
		var item model.Book
		savedItem = savedItems[i]
		item.ID = savedItem.ID
		item.Title = savedItem.Title
		item.Author = savedItem.Author
		items = append(items, &item)
	}
	return items, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
