package usecases

import (
	"errors"

	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/services"
)

type LoginAuthProps struct {
	Email    string
	Password string
}

type loginUseCase struct {
	authRepo repositories.AuthRepository
	userRepo repositories.UserRepository
	jwtService services.IJWTService
}

func NewLoginUseCase(authRepo repositories.AuthRepository, userRepo repositories.UserRepository, jwtService services.IJWTService) *loginUseCase {
	return &loginUseCase{authRepo: authRepo, jwtService: jwtService}
}

func (uc *loginUseCase) Execute(props LoginAuthProps) (*string, error) {
	user, err := uc.userRepo.FindByEmail(props.Email)
	if err != nil {
		return nil, err
	}

	auth, err := uc.authRepo.FindByEmail(props.Email)
	if err != nil {
		return nil, err
	}

	var token *string
	if auth != nil && auth.GetPassword() == props.Password {
		token, err = uc.jwtService.GenerateToken(user.GetID())
		if err != nil {
			return nil, err
		}
	}
	
	return token, errors.New("invalid email or password")
}
