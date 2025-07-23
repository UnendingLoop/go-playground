package migration

import (
	"fmt"
	repka "main/internal/repository"

	"github.com/jmoiron/sqlx"
)

// InitializeDBTable -
func InitializeDBTable() (*repka.DBConn, error) {
	repo := &repka.DBConn{}

	//db, err := sql.Open("sqlite3", ":memory:")
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		return repo, fmt.Errorf("Couldn't connect to DB: %w", err)
	}
	if err = db.Ping(); err != nil {
		return repo, fmt.Errorf("Couldn't connect to DB: %w", err)
	}

	repo = &repka.DBConn{
		DB: db,
	}

	//Creating table
	createTable := `
	CREATE TABLE users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	surname TEXT NOT NULL,
	balance FLOAT NOT NULL,
	currency TEXT NOT NULL
	);`
	db.MustExec(createTable)
	return repo, nil
}
