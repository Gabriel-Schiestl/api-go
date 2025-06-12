package usecases

import (
	"time"

	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type createEventUseCase struct{
	eventRepository repositories.IEventRepository
}

func NewCreateEventUseCase(eventRepository repositories.IEventRepository) *createEventUseCase {
	return &createEventUseCase{
		eventRepository: eventRepository,
	}
}

type CreateEventProps struct {
	Name        string
	Location    string
	Date        time.Time
	Description string
	OrganizerID string
	Category	string
}

func (uc *createEventUseCase) Execute(props CreateEventProps) (*dtos.EventDto, error) {
	var event models.Event

	event, businessErr := models.NewEvent(models.EventProps{
		Name:        &props.Name,
		Location:    &props.Location,
		Date:        &props.Date,
		Description: &props.Description,
		OrganizerID: &props.OrganizerID,
		Category: 	 &props.Category,
	}); 
	if businessErr != nil {
		return nil, businessErr
	}

	err := uc.eventRepository.Save(event)
	if err != nil {
		return nil, err
	}

	return &dtos.EventDto{
		ID:          event.ID(),
		Name:        event.Name(),
		Location:    event.Location(),
		Date:        event.Date(),
		Description: event.Description(),
		OrganizerID: event.OrganizerID(),
		Attendees:   event.Attendees(),
		CreatedAt:   event.CreatedAt(),
		Category: 	 event.Category(),
	}, nil
}