package db

import (
	"github.com/jinzhu/gorm"
	"github.com/protonhq/proton/domain"
	"github.com/protonhq/proton/repository"

	"golang.org/x/crypto/bcrypt"
)

type accountRepository struct {
	repository.AccountRepository
	db *gorm.DB
}

// NewAccountRepository - Create AccountRepositroy
func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(email string, password string) (*domain.Account, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account := domain.Account{Email: email, Password: hash}
	r.db.Create(&account)

	return &account, nil
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

func (r *accountRepository) FindByEmail(email string) (*domain.Account, error) {
	var accts []domain.Account
	r.db.Find(&accts, domain.Account{Email: email})
	if len(accts) != 0 {
		return &accts[0], nil
	}
	return nil, nil
}
