package models

import (
	"time"

	"github.com/Gabriel-Schiestl/go-clarch/domain/exceptions"
	"github.com/google/uuid"
)

type EventProps struct {
    ID          *string
    Name        *string
    Location    *string
    Date        *time.Time
    Description *string
    OrganizerID *string
    Attendees   []string
    CreatedAt   *time.Time
}

type event struct {
    id          string
    name        string
    location    string
    date        time.Time
    description string
    organizerID string
    attendees   []string
    createdAt   time.Time
}

type Event interface {
    ID() string
    Name() string
    Location() string
    Date() time.Time
    Description() string
    OrganizerID() string
    Attendees() []string
    CreatedAt() time.Time
}

func NewEvent(props EventProps) (Event, *exceptions.BusinessException) {
	if props.Name == nil || *props.Name == "" {
		return nil, exceptions.NewBusinessException("Event name is required")
	}
	if props.Location == nil || *props.Location == "" {
		return nil, exceptions.NewBusinessException("Event location is required")
	}
	if props.Date == nil {
		return nil, exceptions.NewBusinessException("Event date is required")
	}
	if props.OrganizerID == nil || *props.OrganizerID == "" {
		return nil, exceptions.NewBusinessException("Organizer ID is required")
	}

	event := &event{
		name:        *props.Name,
		location:    *props.Location,
		date:        *props.Date,
		description: *props.Description,
		organizerID: *props.OrganizerID,
		attendees:   props.Attendees,
		createdAt:   time.Now(),
    }

    if props.ID == nil || *props.ID == "" {
        event.id = uuid.NewString()
    } else {
        event.id = *props.ID
    }

	if props.CreatedAt != nil {
        event.createdAt = *props.CreatedAt
    }

	return event, nil
}

func LoadEvent(props EventProps) (Event, error) {
	return NewEvent(props)
}

func (e *event) ID() string { return e.id }
func (e *event) Name() string { return e.name }
func (e *event) Location() string { return e.location }
func (e *event) Date() time.Time { return e.date }
func (e *event) Description() string { return e.description }
func (e *event) OrganizerID() string { return e.organizerID }
func (e *event) Attendees() []string { return e.attendees }
func (e *event) CreatedAt() time.Time { return e.createdAt }