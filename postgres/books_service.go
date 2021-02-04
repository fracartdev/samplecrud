package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/fracartdev/samplecrud/books"
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

// Create metodo per aggiungere un libro al db
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

// Read metodo per leggere un libro dal db
func (t *BooksService) Read(id string) (*books.BookItem, error) {
	selectSQL := "select id, title, author, createdAt, updatedAt from books where id = $1"
	dbPool := t.getConnection()
	defer dbPool.Close()

	var bookItem books.BookItem
	err := dbPool.QueryRow(context.Background(), selectSQL, id).Scan(&bookItem.ID, &bookItem.Title, &bookItem.Author, &bookItem.CreatedOn, &bookItem.UpdatedOn)
	if err != nil {
		return nil, err
	}
	return &bookItem, nil
}

// Update metodo per aggiornare libro
func (t *BooksService) Update(id string, title string, author string) error {
	updateSQL := "update books set title = $1, author = $2 where id = $3"
	ctx := context.Background()
	dbPool := t.getConnection()
	defer dbPool.Close()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, updateSQL, title, author, id)
	if err != nil {
		log.Println("ERROR: Could not save the Book item due to the error:", err)
		rollbackErr := tx.Rollback(ctx)
		log.Fatal("ERROR: Transaction rollback failed due to the error: ", rollbackErr)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete metodo per cancellare libro
func (t *BooksService) Delete(id string) error {
	deleteSQL := "delete from books where id = $1"

	ctx := context.Background()
	dbPool := t.getConnection()
	defer dbPool.Close()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, deleteSQL, id)
	if err != nil {
		log.Println("ERROR: Could not delete the Book item due to the error:", err)
		rollbackErr := tx.Rollback(ctx)
		log.Fatal("ERROR: Transaction rollback failed due to the error: ", rollbackErr)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
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
