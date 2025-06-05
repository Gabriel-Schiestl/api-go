package repositories

import "github.com/Gabriel-Schiestl/api-go/internal/domain/models"

type UserRepository interface {
	Create(user models.User) error
	FindAll() ([]models.User, error)
}
