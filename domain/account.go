package domain

import "github.com/jinzhu/gorm"

// Account - Account entity
type Account struct {
	gorm.Model
	Email string `json:"email"`
}
