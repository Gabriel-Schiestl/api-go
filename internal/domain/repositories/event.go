package repositories

import (
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
)

type IRepository interface {
	FindByID(id string) (models.Event, error)
	FindAll() ([]models.Event, error)
	Save(event models.Event) error
	Delete(id string) error
}