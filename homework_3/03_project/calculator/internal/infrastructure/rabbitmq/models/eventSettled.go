package models

// EventSettled represents a DTO for settled events.
type EventSettled struct {
	Id      string `json:"id"`
	Outcome string `json:"outcome"`
}
