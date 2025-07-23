package repository

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func createTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatal("Couldn't connect to DB: %w", err)
	}
	//Ping не требуется так как SQLX сразу пингует и выдает ошибку подключения
	// if err = db.Ping(); err != nil {
	// 	t.Fatal("Couldn't connect to DB: %w", err)
	// }

	//Creating table
	createTable := `
	CREATE TABLE users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	surname TEXT NOT NULL,
	balance INTEGER NOT NULL,
	currency TEXT NOT NULL
	);`
	if _, err := db.Exec(createTable); err != nil {
		t.Fatal("Couldn't create table: %w", err)
	}
	return db
}
func addMockData(t *testing.T, db *sqlx.DB) {
	tx, err := db.Begin()
	if err != nil {
		t.Fatal("Transaction start error: %w", err)
	}
	defer tx.Rollback()

}
func TestCreateUser(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	testStruct := DBConn{DB: db}
	tx, err := db.Beginx()
	if err != nil {
		t.Fatal("Transaction start error: %w", err)
	}
	defer tx.Rollback()

	err = testStruct.CreateUser(tx, "Name", "Surname", "EUR", 1000)
	if err != nil {
		tx.Rollback()
		t.Fatal("Error creating user: %w", err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal("Commit error: %w", err)
	}

	row := db.QueryRow(`SELECT * FROM users WHERE name=? AND surname=?`, "Name", "Surname")
	var id, balance int
	var name, surname, currency string
	if err := row.Scan(&id, &name, &surname, &balance, &currency); err != nil {
		t.Fatal("Error reading row: %w", err)
	}
	if name != "Name" {
		t.Fatalf("Expected: 'Name', got: %s", name)
	}
	if surname != "Surname" {
		t.Fatalf("Expected: 'Surname', got: %s", name)
	}
	if balance != 1000 {
		t.Fatalf("Expected: 1000, got: %d", balance)
	}
	if currency != "EUR" {
		t.Fatalf("Expected: 'EUR', got: %s", currency)
	}
}

func TestBalanceOp(t *testing.T) {
	db := createTestDB(t)
	addMockData(t, db)

}
