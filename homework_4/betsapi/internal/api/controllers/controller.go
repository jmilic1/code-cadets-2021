package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const idString = "id"
const statusString = "status"
const statusDefaultValue = "active"

// Controller implements handlers for web server requests.
type Controller struct {
	betService BetService
}

// NewController creates a new instance of Controller
func NewController(betService BetService) *Controller {
	return &Controller{
		betService: betService,
	}
}

// GetBet handles get bet request.
func (c *Controller) GetBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		betId := ctx.Param(idString)

		dto, found, err := c.betService.GetBet(ctx, betId)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}
		if !found {
			ctx.String(http.StatusNotFound, "bet with given id does not exist.")
			return
		}

		ctx.JSON(200, dto)
	}
}

// GetBetsByCustomerId handles get bet by customerId request.
func (c *Controller) GetBetsByCustomerId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerId := ctx.Param(idString)
		dtoModels, err := c.betService.GetBetsByCustomerId(ctx, customerId)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}
		if len(dtoModels) == 0 {
			ctx.String(http.StatusNotFound, "given customer id has no attributed bets.")
			return
		}

		ctx.JSON(200, dtoModels)
	}
}

// GetBetsByStatus handles get bet by status request
func (c *Controller) GetBetsByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.DefaultQuery(statusString, statusDefaultValue)

		dtoModels, err := c.betService.GetBetsByStatus(ctx, status)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}
		if len(dtoModels) == 0 {
			ctx.String(http.StatusNotFound, "no bet has given status")
			return
		}

		ctx.JSON(200, dtoModels)
	}
}
