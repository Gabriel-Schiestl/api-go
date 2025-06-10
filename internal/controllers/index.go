package controllers

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/usecases"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/database"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/database/connection"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/mappers"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/ports"
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
	
	userMapper := mappers.UserMapper{}
	userRepository := database.NewUserRepository(connection.Db, userMapper)

	getUsersUseCase := usecases.NewGetUsersUseCase(userRepository)
	getUsersDecorator := usecase.NewUseCaseDecorator(getUsersUseCase)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepository)
	createUserDecorator := usecase.NewUseCaseWithPropsDecorator(createUserUseCase)

	usersController := NewUsersController(createUserDecorator, getUsersDecorator)
	controller.Add(usersController)

	authMapper := mappers.AuthMapper{}
	authRepository := database.NewAuthRepository(connection.Db, authMapper)
	jwtService := ports.NewJWTService()


	getAuthsUseCase := usecases.NewGetAuthsUseCase(authRepository)
	getAuthsDecorator := usecase.NewUseCaseDecorator(getAuthsUseCase)
	createAuthUseCase := usecases.NewCreateAuthUseCase(authRepository)
	createAuthDecorator := usecase.NewUseCaseWithPropsDecorator(createAuthUseCase)
	loginUseCase := usecases.NewLoginUseCase(authRepository, userRepository, jwtService)
	loginDecorator := usecase.NewUseCaseWithPropsDecorator(loginUseCase)

	authController := NewAuthController(createAuthDecorator, getAuthsDecorator, loginDecorator)
	controller.Add(authController)
}