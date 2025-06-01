package dtos

type EventDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Location    string   `json:"location"`
	Date        string   `json:"date"`
	Description string   `json:"description"`
	OrganizerID string   `json:"organizer_id"`
	Attendees   []string `json:"attendees"`
	CreatedAt   string   `json:"created_at"`
}