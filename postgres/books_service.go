package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
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

	conn, err := sql.Open("pgx", dbConnectionString)

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
	insertSQL := "insert into books(id, title, author) values ($1, $2, $3)"

	ctx := context.Background()
	dbPool := t.getConnection()
	defer dbPool.Close()

	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	id := uuid.New()
	idStr := id.String()

	_, err = tx.Exec(ctx, insertSQL, idStr, title, author)
	if err != nil {
		log.Println("ERROR: Could not save the Book item due to the error:", err)
		rollbackErr := tx.Rollback(ctx)
		log.Fatal("ERROR: Transaction rollback failed due to the error: ", rollbackErr)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &idStr, nil
}

func (t *BooksService) getConnection() *pgxpool.Pool {
	dbPool, err := pgxpool.Connect(context.Background(), t.getDBConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	return dbPool
}

func (t *BooksService) getDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", t.DbUserName, t.DbPassword, t.DbHost, t.DbName)
}
