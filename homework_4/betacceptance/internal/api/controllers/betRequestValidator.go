package controllers

import "github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/api/controllers/models"

// BetRequestValidator validates bet requests.
type BetRequestValidator interface {
	BetRequestIsValid(betRequestDto models.BetRequestDto) bool
}
