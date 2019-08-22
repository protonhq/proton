package usecase

import (
	"errors"

	"github.com/protonhq/proton/domain"
	"github.com/protonhq/proton/repository"
)

// AccountUsecase - account usecase
type AccountUsecase interface {
	RegisterUser(email string, password string) (*domain.Account, error)
}

type accountUsecase struct {
	AccountUsecase
	repo repository.AccountRepository
}

// NewAccountUsecase - Create account usecase
func NewAccountUsecase(repo repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		repo: repo,
	}
}

// RegisterUser -  register a user
func (a *accountUsecase) RegisterUser(email string, password string) (*domain.Account, error) {
	acct, err := a.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if acct != nil {
		return nil, errors.New("Email already exist")
	}

	newAcct, err := a.repo.Create(email, password)

	if err != nil {
		return nil, err
	}

	return newAcct, nil
}
