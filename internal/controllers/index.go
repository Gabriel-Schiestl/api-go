package controllers

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/usecases"
	"github.com/Gabriel-Schiestl/go-clarch/application/usecase"
	"github.com/Gabriel-Schiestl/go-clarch/presentation/controller"
)

var Controllers = []controller.Controller{}

func init() {
	getEventsUseCase := usecases.NewGetEventsUseCase()
	getEventsDecorator := usecase.NewUseCaseDecorator(getEventsUseCase)
	eventsController := NewEventsController(getEventsDecorator)
	controller.Add(eventsController)
}