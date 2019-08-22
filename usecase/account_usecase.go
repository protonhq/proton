package usecase

import (
	"errors"

	"github.com/protonhq/proton/repository"
)

// AccountUsecase - account usecase
type AccountUsecase interface {
	RegisterUser(email string, password string) error
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
func (a *accountUsecase) RegisterUser(email string, password string) error {
	acct, err := a.repo.FindByEmail(email)
	if err != nil {
		return err
	}

	if acct != nil {
		return errors.New("Email already exist")
	}

	return a.repo.Create(email, password)
}
