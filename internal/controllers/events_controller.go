package controllers

import (
	r "github.com/Gabriel-Schiestl/api-go/internal/server"
	"github.com/Gabriel-Schiestl/go-clarch/application/usecase"
	_ "github.com/Gabriel-Schiestl/go-clarch/presentation/controller"
	"github.com/gin-gonic/gin"
)

type EventsController[R any] struct{
	getEventsUseCase usecase.UseCaseDecorator[R]
}

func NewEventsController[R any](getEventsUseCase usecase.UseCaseDecorator[R]) *EventsController[R] {
	return &EventsController[R]{
		getEventsUseCase: getEventsUseCase,
	}
}

func (ec EventsController[R]) GetAllEvents(c *gin.Context) {
	events, err := ec.getEventsUseCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, events)
}

func (ec EventsController[R]) SetupRoutes() {
	group := r.Router.Group("/events")

	group.GET("/", ec.GetAllEvents)
}