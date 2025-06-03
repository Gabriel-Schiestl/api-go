package dtos

import "time"

type EventDto struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	OrganizerID string    `json:"organizer_id"`
	Attendees   []string  `json:"attendees"`
	CreatedAt   time.Time    `json:"created_at"`
}