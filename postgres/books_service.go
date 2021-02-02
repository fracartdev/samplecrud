package postgres

import (
	"database/sql"
	"fmt"
)

// BooksService struct mi mette a disposizione le dipendenze di cui ho ibisogno per implementare l'interfaccia
type BooksService struct {
	DbUserName string
	DbPassword string
	DbHost     string
	DbName     string
	DbPort     int
}

// Init inizializza la connessione al db
func (t *BooksService) Init() error {
	dbConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", t.DbHost, t.DbPort, t.DbUserName, t.DbPassword, t.DbName)

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("DB Connesso!")
	return nil
}

// Create metodo per implementare l'interfaccia BookService
func (t *BooksService) Create(title string, author string) (*string, error) {
	panic("Da implementare")
}
