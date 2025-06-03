package controllers

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/application/usecases"
	r "github.com/Gabriel-Schiestl/api-go/internal/server"
	"github.com/Gabriel-Schiestl/go-clarch/application/usecase"
	_ "github.com/Gabriel-Schiestl/go-clarch/presentation/controller"
	"github.com/gin-gonic/gin"
)

type EventsController struct{
	getEventsUseCase usecase.UseCaseDecorator[[]dtos.EventDto]
	createEventUseCase usecase.UseCaseWithPropsDecorator[usecases.CreateEventProps, *dtos.EventDto]
}

func NewEventsController(getEventsUseCase usecase.UseCaseDecorator[[]dtos.EventDto], createEventUseCase usecase.UseCaseWithPropsDecorator[usecases.CreateEventProps, *dtos.EventDto]) *EventsController {
	return &EventsController{
		getEventsUseCase: getEventsUseCase,
		createEventUseCase: createEventUseCase,
	}
}

func (ec EventsController) GetAllEvents(c *gin.Context) {
	events, err := ec.getEventsUseCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, events)
}

func (ec EventsController) CreateEvent(c *gin.Context) {
	body := usecases.CreateEventProps{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	event, err := ec.createEventUseCase.Execute(body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, event)
}

func (ec EventsController) SetupRoutes() {
	group := r.Router.Group("/events")

	group.GET("/", ec.GetAllEvents)
	group.POST("/", ec.CreateEvent)
}