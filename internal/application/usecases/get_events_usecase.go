package usecases

import (
	"time"

	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
)

type getEventsUseCase struct{}

func NewGetEventsUseCase() *getEventsUseCase {
	return &getEventsUseCase{}
}

func (uc *getEventsUseCase) Execute() ([]dtos.EventDto, error) {
	name := "Test Event"
    location := "Test Location"
    date := "2025-06-15"
    description := "Test Description"
    organizerID := "org-123"

    event, _ := models.NewEvent(models.EventProps{
        ID:          nil,
        Name:        &name,
        Location:    &location,
        Date:        &date,
        Description: &description,
        OrganizerID: &organizerID,
        Attendees:   []string{"user-1", "user-2"},
        CreatedAt:   nil,
    })

	return []dtos.EventDto{
		dtos.EventDto{
			ID:          event.ID(),
			Name:        event.Name(),
			Location:    event.Location(),
			Date:        event.Date(),
			Description: event.Description(),
			OrganizerID: event.OrganizerID(),
			Attendees:   event.Attendees(),
			CreatedAt:   event.CreatedAt().Format(time.RFC3339),
		},
	}, nil
}
