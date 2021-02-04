package books

import "time"

// BookItem oggetto per gestire i libri nelle funzioni CRUD
type BookItem struct {
	ID        string
	Title     string
	Author    bool
	CreatedOn time.Time
	UpdatedOn *time.Time
}

// Book interfaccia
type Book interface {
	Init() error
	Create(title string, author string) (*string, error)
	Read(id string) (*BookItem, error)
}
