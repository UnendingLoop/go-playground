package services

import (
	"fmt"
	repka "main/internal/repository"

	"github.com/jmoiron/sqlx"
)

// UserService - struct with interface
type UserService struct {
	repo repka.Repository
}

// NewUserService - creates a new instance from a structure that meets the interface UserService
func NewUserService(r repka.Repository) UserService {
	return UserService{repo: r}
}

// CreateUser - creates a new user in DB
func (s *UserService) CreateUser(tx *sqlx.Tx, name, surname, currency string, balance int) error {
	if err := s.repo.CreateUser(tx, name, surname, currency, balance); err != nil {
		tx.Rollback()
		return fmt.Errorf("Couldn't finish operation: %w", err)
	}
	return nil
}

// AlterBalance -
func (s *UserService) AlterBalance(tx *sqlx.Tx, name, surname string, balanceDelta int) error {
	if err := s.repo.AlterBalance(tx, name, surname, balanceDelta); err != nil {
		tx.Rollback()
		return fmt.Errorf("Couldn't finish operation: %w", err)
	}
	return nil
}

// GetUserInfo -
func (s *UserService) GetUserInfo(name, surname string) (repka.Info, error) {
	return s.repo.GetUserInfo(name, surname)
}
