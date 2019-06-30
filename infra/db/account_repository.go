package db

import (
	"github.com/jinzhu/gorm"
	"github.com/protonhq/proton/domain"
	"github.com/protonhq/proton/domain/repository"
)

type accountRepository struct {
	repository.AccountRepository
	db *gorm.DB
}

// NewAccountRepository - Create AccountRepositroy
func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(email string, password string) error {
	return nil
}

func (r *accountRepository) Get(id string) (*domain.Account, error) {
	return nil, nil
}

func (r *accountRepository) GetAll([]*domain.Account, error) {
	return
}

func (r *accountRepository) Save(*domain.Account) error {
	return nil
}

func (r *accountRepository) Remove(id string) error {
	return nil
}

func (r *accountRepository) Update(*domain.Account) error {
	return nil
}
