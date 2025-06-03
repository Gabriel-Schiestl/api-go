package database

import (
	"fmt"

	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/entities"
	"github.com/Gabriel-Schiestl/go-clarch/domain/exceptions"
	"gorm.io/gorm"
)

type eventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) repositories.IEventRepository {
	return eventRepositoryImpl{
		db: db,
	}
}

func (r eventRepositoryImpl) FindByID(id string) (models.Event, error) {
	var event entities.Event
	if err := r.db.First(&event, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRepositoryNoDataFoundException(fmt.Sprintf("Event with ID %s not found", id))
		}

		return nil, exceptions.NewTechnicalException(fmt.Sprintf("Error retrieving event with ID %s: %v", id, err))
	}

	domain, err := models.LoadEvent(models.EventProps{
		ID:          &event.ID,
		Name:        &event.Name,
		Location:    &event.Location,
		Date:        &event.Date,
		Description: &event.Description,
		OrganizerID: &event.OrganizerID,
		Attendees:   event.Attendees,
		CreatedAt:   &event.CreatedAt,
	})
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r eventRepositoryImpl) FindAll() ([]models.Event, error) {
	var events []entities.Event
	if err := r.db.Find(&events).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRepositoryNoDataFoundException("No events found")
		}

		return nil, exceptions.NewTechnicalException(fmt.Sprintf("Error retrieving events: %v", err))
	}

	var domainEvents []models.Event
	for _, event := range events {
		domain, err := models.LoadEvent(models.EventProps{
			ID:          &event.ID,
			Name:        &event.Name,
			Location:    &event.Location,
			Date:        &event.Date,
			Description: &event.Description,
			OrganizerID: &event.OrganizerID,
			Attendees:   event.Attendees,
			CreatedAt:   &event.CreatedAt,
		})
		if err != nil {
			return nil, err
		}

		domainEvents = append(domainEvents, domain)
	}

	return domainEvents, nil
}

func (r eventRepositoryImpl) Save(event models.Event) error {
	entity := entities.Event{
		ID: 		event.ID(),
		Name:        event.Name(),
		Location:    event.Location(),
		Date:        event.Date(),
		Description: event.Description(),
		OrganizerID: event.OrganizerID(),
		Attendees:   event.Attendees(),
		CreatedAt:   event.CreatedAt(),
	}
	if err := r.db.Save(&entity).Error; err != nil {
		return exceptions.NewTechnicalException(fmt.Sprintf("Error saving event: %v", err))
	}

	return nil
}

func (r eventRepositoryImpl) Delete(id string) error {
	var event entities.Event
	if err := r.db.First(&event, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewRepositoryNoDataFoundException(fmt.Sprintf("Event with ID %s not found", id))
		}
		return exceptions.NewTechnicalException(fmt.Sprintf("Error retrieving event with ID %s: %v", id, err))
	}

	if err := r.db.Delete(&event).Error; err != nil {
		return exceptions.NewTechnicalException(fmt.Sprintf("Error deleting event with ID %s: %v", id, err))
	}

	return nil
}
