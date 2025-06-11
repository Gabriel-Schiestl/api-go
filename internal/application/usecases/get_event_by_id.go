package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type getEventByIdUseCase struct {
	eventRepo repositories.IEventRepository
}

func NewGetEventByIdUseCase(eventRepo repositories.IEventRepository) *getEventByIdUseCase {
	return &getEventByIdUseCase{
		eventRepo: eventRepo,
	}
}

func (uc *getEventByIdUseCase) Execute(id string) (dtos.EventDto, error) {
	event, err := uc.eventRepo.FindByID(id)
	if err != nil {
		return dtos.EventDto{}, err
	}

	eventDto := dtos.EventDto{
		ID:          event.ID(),
		Name:        event.Name(),
		Description: event.Description(),
		Location:    event.Location(),
		Date:        event.Date(),
		OrganizerID: event.OrganizerID(),
		Attendees:   event.Attendees(),
		CreatedAt:   event.CreatedAt(),
	}

	
	return eventDto, nil
}
