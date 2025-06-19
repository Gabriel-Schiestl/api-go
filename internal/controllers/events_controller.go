package controllers

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/application/usecases"
	r "github.com/Gabriel-Schiestl/api-go/internal/server"
	"github.com/Gabriel-Schiestl/go-clarch/application/usecase"
	_ "github.com/Gabriel-Schiestl/go-clarch/presentation/controller"
	"github.com/gin-gonic/gin"
)

var userIDRequired = gin.H{"error": "User ID is required"}
var eventIDRequired = gin.H{"error": "Event ID is required"}

type EventsController struct{
	getEventsUseCase usecase.UseCaseDecorator[[]dtos.EventDto]
	createEventUseCase usecase.UseCaseWithPropsDecorator[usecases.CreateEventProps, *dtos.EventDto]
	getEventsByUserUseCase usecase.UseCaseWithPropsDecorator[string, []dtos.EventDto]
	getEventByIdUseCase usecase.UseCaseWithPropsDecorator[string, dtos.EventDto]
	registerToEventUseCase usecase.UseCaseWithPropsDecorator[usecases.RegisterToEventUseCaseProps, []string]
	getEventByOrganizerUseCase usecase.UseCaseWithPropsDecorator[usecases.GetEventByOrganizerUseCaseProps, dtos.EventWithAttendeesDto]
	getEventsByOrganizerUseCase usecase.UseCaseWithPropsDecorator[string, []dtos.EventDto]
}

func NewEventsController(
	getEventsUseCase usecase.UseCaseDecorator[[]dtos.EventDto],
	createEventUseCase usecase.UseCaseWithPropsDecorator[usecases.CreateEventProps, *dtos.EventDto],
	getEventsByUserUseCase usecase.UseCaseWithPropsDecorator[string, []dtos.EventDto],
	getEventByIdUsecase usecase.UseCaseWithPropsDecorator[string, dtos.EventDto],
	registerToEventUseCase usecase.UseCaseWithPropsDecorator[usecases.RegisterToEventUseCaseProps, []string],
	getEventByOrganizerUseCase usecase.UseCaseWithPropsDecorator[usecases.GetEventByOrganizerUseCaseProps, dtos.EventWithAttendeesDto],
	getEventsByOrganizerUseCase usecase.UseCaseWithPropsDecorator[string, []dtos.EventDto],
) *EventsController {
	return &EventsController{
		getEventsUseCase: getEventsUseCase,
		createEventUseCase: createEventUseCase,
		getEventsByUserUseCase: getEventsByUserUseCase,
		getEventByIdUseCase: getEventByIdUsecase,
		registerToEventUseCase: registerToEventUseCase,
		getEventByOrganizerUseCase: getEventByOrganizerUseCase,
		getEventsByOrganizerUseCase: getEventsByOrganizerUseCase,
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
	userID, exists := c.Get("userID")
	if !exists || userID == "" {
		c.JSON(400, userIDRequired)
		return
	}

	body := usecases.CreateEventProps{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	body.OrganizerID = userID.(string)

	_, err := ec.createEventUseCase.Execute(body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Event created successfully"})
}

func (ec EventsController) GetEventsByUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists || userID == "" {
		c.JSON(400, userIDRequired)
		return
	}

	events, err := ec.getEventsByUserUseCase.Execute(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, events)
}

func (ec EventsController) GetEventById(c *gin.Context) {
	eventID := c.Param("eventID")
	if eventID == "" {
		c.JSON(400, eventIDRequired)
		return
	}

	event, err := ec.getEventByIdUseCase.Execute(eventID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, event)
}

func (ec EventsController) RegisterToEvent(c *gin.Context) {
	eventID := c.Param("eventID")
	userID, exists := c.Get("userID")
	if !exists || userID == "" {
		c.JSON(400, userIDRequired)
		return
	}

	if eventID == "" {
		c.JSON(400, eventIDRequired)
		return
	}

	props := usecases.RegisterToEventUseCaseProps{
		UserId: userID.(string),
		EventId: eventID,
	}

	attendees, err := ec.registerToEventUseCase.Execute(props)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, attendees)
}

func (ec EventsController) GetEventByOrganizer(c *gin.Context) {
	eventId := c.Param("eventId")
	userID, exists := c.Get("userID")
	if !exists || userID == "" {
		c.JSON(400, userIDRequired)
		return
	}

	if eventId == "" {
		c.JSON(400, eventIDRequired)
		return
	}

	props := usecases.GetEventByOrganizerUseCaseProps{
		OrganizerId: userID.(string),
		EventId:     eventId,
	}

	event, err := ec.getEventByOrganizerUseCase.Execute(props)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, event)
}

func (ec EventsController) GetEventsByOrganizer(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists || userID == "" {
		c.JSON(400, userIDRequired)
		return
	}

	events, err := ec.getEventsByOrganizerUseCase.Execute(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, events)
}

func (ec EventsController) SetupRoutes() {
	group := r.Router.Group("/events")

	group.GET("/", ec.GetAllEvents)
	group.POST("/", ec.CreateEvent)
	group.GET("/registered", ec.GetEventsByUser)
	group.GET("/:eventID", ec.GetEventById)
	group.POST("/:eventID/register", ec.RegisterToEvent)
	group.GET("/:eventID/organizer", ec.GetEventByOrganizer)
	group.GET("/organizer", ec.GetEventsByOrganizer)
}