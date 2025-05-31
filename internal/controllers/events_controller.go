package controllers

import (
	r "github.com/Gabriel-Schiestl/api-go/internal/server"
	"github.com/Gabriel-Schiestl/api-go/pkg"
	"github.com/gin-gonic/gin"
)

type EventsController struct{
	getEventsUseCase pkg.UseCaseDecorator
}

func NewEventsController(getEventsUseCase pkg.UseCaseDecorator) *EventsController {
	return &EventsController{
		getEventsUseCase: getEventsUseCase,
	}
}

func (ec EventsController) GetAllEvents(c *gin.Context) {
	events, err := ec.getEventsUseCase.Execute(nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, events)
}

func (ec EventsController) SetupRoutes() {
	group := r.Router.Group("/events")

	group.GET("/", ec.GetAllEvents)
}