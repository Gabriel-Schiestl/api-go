package controllers

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/usecases"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/database"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/database/connection"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/mappers"
	"github.com/Gabriel-Schiestl/go-clarch/application/usecase"
	"github.com/Gabriel-Schiestl/go-clarch/presentation/controller"
)

var Controllers = []controller.Controller{}

func SetupControllers() {
	mapper := mappers.EventMapper{}
	eventRepository := database.NewEventRepository(connection.Db, mapper)

	getEventsUseCase := usecases.NewGetEventsUseCase(eventRepository)
	getEventsDecorator := usecase.NewUseCaseDecorator(getEventsUseCase)

	createEventUseCase := usecases.NewCreateEventUseCase(eventRepository)
	createEventDecorator := usecase.NewUseCaseWithPropsDecorator(createEventUseCase)

	eventsController := NewEventsController(getEventsDecorator, createEventDecorator)
	controller.Add(eventsController)
}