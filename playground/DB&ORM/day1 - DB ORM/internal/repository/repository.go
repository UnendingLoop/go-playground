package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Repository - interface with methods to work with DB
type Repository interface {
	CreateUser(tx *sqlx.Tx, name, surname, currency string, balance int) error
	AlterBalance(tx *sqlx.Tx, name, surname string, balanceDelta int) error
	GetUserInfo(name, surname string) (Info, error)
}

// DBConn - structure to contain a db-connection
type DBConn struct {
	DB *sqlx.DB
}

// Info -
type Info struct {
	ID       int     `db:"id"`
	Name     string  `db:"name"`
	Surname  string  `db:"surname"`
	Currency string  `db:"currency"`
	Balance  float64 `db:"balance"`
}

// CreateUser -
func (r *DBConn) CreateUser(tx *sqlx.Tx, name, surname, currency string, balance int) error {
	_, err := tx.Exec(`INSERT INTO users (name, surname, balance, currency) VALUES (?,?,?,?)`, name, surname, balance, currency)
	return err
}

// AlterBalance -
func (r *DBConn) AlterBalance(tx *sqlx.Tx, name, surname string, balanceDelta int) error {
	row := tx.QueryRowx(`SELECT COUNT(*) FROM users WHERE name = ? AND surname = ?`, name, surname)
	count := 0
	if err := row.Scan(&count); err != nil {
		return fmt.Errorf("Couldn't read userdata to alter balance: %s", err)
	}
	if count > 1 || count == 0 {
		return fmt.Errorf("Cannot alter balance: found %d of records that meet search requirements", count)
	}

	row = tx.QueryRowx(`SELECT id, balance FROM users WHERE name = ? AND surname = ?`, name, surname)
	id := 0
	balance := 0
	if err := row.Scan(&id, &balance); err != nil {
		return fmt.Errorf("Couldn't read userdata to alter balance: %s", err)
	}
	if balance+balanceDelta < 0 {
		return fmt.Errorf("Insufficient funds to alter balance: current balance is %d, delta is %d", balance, balanceDelta)
	}

	_, err := tx.Exec(`UPDATE users SET balance=? WHERE id=?`, balance+balanceDelta, id)
	if err != nil {
		return fmt.Errorf("Couldn't update balance: %s", err)
	}
	return nil
}

// GetUserInfo -
func (r *DBConn) GetUserInfo(name, surname string) (Info, error) {
	user := Info{
		Name:    name,
		Surname: surname,
	}
	err := r.DB.Get(&user, `SELECT id, currency, balance FROM users WHERE name = ? AND surname = ?`, name, surname)
	if err != nil {
		return Info{}, fmt.Errorf("Error fetching userinfo: %s", err)
	}
	return user, nil
}
