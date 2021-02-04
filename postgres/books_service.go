package postgres

import (
	"database/sql"
	"fmt"

	// importo il driver
	_ "github.com/lib/pq"
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

	conn, err := sql.Open("postgres", dbConnectionString)

	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	fmt.Println("DB Connesso!")
	return nil
}

// Create metodo per implementare l'interfaccia BookService
func (t *BooksService) Create(title string, author string) (*string, error) {
	var id int
	var createdAt string
	query := `INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, item.Name, item.Description).Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	item.ID = id
	item.CreatedAt = createdAt
	return nil
}
