package repositories

import "github.com/Gabriel-Schiestl/api-go/internal/domain/models"

type AuthRepository interface {
	Create(auth models.Auth) error
	FindAll() ([]models.Auth, error)
}
