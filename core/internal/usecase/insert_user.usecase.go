package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-back/core/internal/entity"
	"github.com/beriloqueiroz/desafio-back/core/internal/usecase/interfaces"
	"github.com/google/uuid"
)

type InsertUserUseCase struct {
	UserRepository interfaces.UserRepository
}

func (u *InsertUserUseCase) Execute(ctx context.Context, email string, phone string, city string, state string) (string, error) {
	location := entity.Location{
		City:  city,
		State: state,
	}
	user, err := entity.NewUser(uuid.NewString(), true, email, phone, location)
	if err != nil {
		return "", err
	}
	err = u.UserRepository.Save(ctx, user)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}
