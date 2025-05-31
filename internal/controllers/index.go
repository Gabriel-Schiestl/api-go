package controllers

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/usecases"
	"github.com/Gabriel-Schiestl/api-go/pkg"
)

var Controllers = []Controller{}

func Init() {
	getEventsUseCase := usecases.NewGetEventsUseCase()
	getEventsDecorator := pkg.NewUseCaseDecorator(getEventsUseCase)
	eventsController := NewEventsController(getEventsDecorator)
	Controllers = append(Controllers, eventsController)

	setupRoutes()
}

func setupRoutes() {
	for _, controller := range Controllers {
		controller.SetupRoutes()
	}
}