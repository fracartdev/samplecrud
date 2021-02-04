package books

import "time"

// BookItem oggetto per gestire i libri nelle funzioni CRUD
type BookItem struct {
	ID        string
	Title     string
	Author    string
	CreatedOn time.Time
	UpdatedOn *time.Time
}

// Book interfaccia
type Book interface {
	Init() error
	Create(title string, author string) (*string, error)
	Read(id string) (*BookItem, error)
	Update(id string, title string, author string) error
	Delete(id string) error
}
