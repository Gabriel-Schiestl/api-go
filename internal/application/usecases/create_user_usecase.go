package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type createUserUseCase struct {
	repo repositories.UserRepository
}

func NewCreateUserUseCase(repo repositories.UserRepository) *createUserUseCase {
	return &createUserUseCase{repo: repo}
}

type CreateUserProps struct {
	Name  string
	Email string
}

func (uc *createUserUseCase) Execute(props CreateUserProps) (*dtos.UserResponseDTO, error) {
	user := models.NewUser(models.UserProps{
		Name:  &props.Name,
		Email: &props.Email,
	})
	err := uc.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return &dtos.UserResponseDTO{
		ID:        user.GetID(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
