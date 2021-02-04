package graph

import "github.com/fracartdev/samplecrud/books"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	BooksLib books.Book
}
