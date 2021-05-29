package models

// BetResultDto bet Get result dto model.
type BetResultDto struct {
	Id                   string  `json:"id"`
	Status               string  `json:"status"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
	Payout               float64 `json:"payout"`
}
