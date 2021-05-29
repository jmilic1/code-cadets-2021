package models

// BetQueueDto represents a DTO of a received bet which should be send to queue.
type BetQueueDto struct {
	Id                   string  `json:"id"`
	CustomerId           string  `json:"customerId"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
}
