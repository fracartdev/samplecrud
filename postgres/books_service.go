package postgres

import "fmt"

// BooksService struct mi mette a disposizione le dipendenze di cui ho ibisogno per implementare l'interfaccia
type BooksService struct {
	Message string
}

// Create metodo per implementare l'interfaccia BookService
func (t *BooksService) Create(title string, author string) (*string, error) {
	fmt.Printf("I have a message: %s\n", t.Message)
	panic("oops!")
}
