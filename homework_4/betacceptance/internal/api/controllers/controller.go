package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	requests "github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/api/controllers/models"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betRequestValidator BetRequestValidator
	betService          BetService
}

// NewController creates a new instance of Controller
func NewController(betRequestValidator BetRequestValidator, betService BetService) *Controller {
	return &Controller{
		betRequestValidator: betRequestValidator,
		betService:          betService,
	}
}

// PostBet handles post bet request.
func (c *Controller) PostBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var betRequest requests.BetRequestDto
		err := ctx.ShouldBindJSON(&betRequest)
		if err != nil {
			ctx.String(http.StatusBadRequest, "update request is not valid.")
			return
		}

		if !c.betRequestValidator.BetRequestIsValid(betRequest) {
			ctx.String(http.StatusBadRequest, "update request is not valid.")
			return
		}

		err = c.betService.SendBet(betRequest)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.Status(http.StatusOK)
	}
}
