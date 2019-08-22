package usecase

import "github.com/protonhq/proton/repository"

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
	return a.repo.Create(email, password)
}
