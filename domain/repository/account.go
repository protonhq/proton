package repository

import "github.com/protonhq/proton/domain"

// AccountRepository - Repository Interface for `Account`
type AccountRepository interface {
	Create(email string, password string) error
	Get(id string) (*domain.Account, error)
	GetAll([]*domain.Account, error)
	Save(*domain.Account) error
	Remove(id string) error
	Update(*domain.Account) error
}
