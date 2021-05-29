package models

// Bet is a storage model representation of a calculated bet.
type Bet struct {
	Id                   string
	SelectionId          string
	SelectionCoefficient int
	Payment              int
}
