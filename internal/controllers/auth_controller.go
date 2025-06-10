package controllers

import (
	"net/http"

	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/application/usecases"
	"github.com/Gabriel-Schiestl/go-clarch/application/usecase"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	createAuthUseCase usecase.UseCaseWithPropsDecorator[usecases.CreateAuthProps, *dtos.AuthResponseDTO]
	getAuthsUseCase   usecase.UseCaseDecorator[[]dtos.AuthResponseDTO]
	loginUseCase  usecase.UseCaseWithPropsDecorator[usecases.LoginAuthProps, *string]
}

func NewAuthController(createUC usecase.UseCaseWithPropsDecorator[usecases.CreateAuthProps, *dtos.AuthResponseDTO], getUC usecase.UseCaseDecorator[[]dtos.AuthResponseDTO], loginUC usecase.UseCaseWithPropsDecorator[usecases.LoginAuthProps, *string]) *AuthController {
	return &AuthController{
		createAuthUseCase: createUC,
		getAuthsUseCase:   getUC,
		loginUseCase:  loginUC,
	}
}

func (c *AuthController) RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("", c.CreateAuth)
		auth.GET("", c.GetAuths)
		auth.POST("/login", c.Login)
	}
}

func (c *AuthController) CreateAuth(ctx *gin.Context) {
	var input usecases.CreateAuthProps
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dto, err := c.createAuthUseCase.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dto)
}

func (c *AuthController) GetAuths(ctx *gin.Context) {
	dtos, err := c.getAuthsUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dtos)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input usecases.LoginAuthProps
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.loginUseCase.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", *token, 3600, "/", "", false, true)

	ctx.JSON(http.StatusOK, token)
}

func (c *AuthController) SetupRoutes() {
	// Use o router global, igual aos outros controllers
	// Exemplo:
	// r := server.Router
	// c.RegisterRoutes(r)
}
