package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type CreateAuthProps struct {
	Email    string
	Password string
}

type createAuthUseCase struct {
	repo repositories.AuthRepository
}

func NewCreateAuthUseCase(repo repositories.AuthRepository) *createAuthUseCase {
	return &createAuthUseCase{repo: repo}
}

func (uc *createAuthUseCase) Execute(props CreateAuthProps) (*dtos.AuthResponseDTO, error) {
	auth := models.NewAuth(models.AuthProps{
		Email:    &props.Email,
		Password: &props.Password,
	})
	err := uc.repo.Create(auth)
	if err != nil {
		return nil, err
	}
	return &dtos.AuthResponseDTO{
		ID:        auth.GetID(),
		Email:     auth.GetEmail(),
		CreatedAt: auth.GetCreatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
