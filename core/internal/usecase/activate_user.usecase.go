package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-back/core/internal/usecase/interfaces"
)

type ActivateUserUseCase struct {
	UserRepository interfaces.UserRepository
}

func (u *ActivateUserUseCase) Execute(ctx context.Context, ID string) error {
	user, err := u.UserRepository.Find(ctx, ID)
	if err != nil {
		return err
	}
	user.Activate()
	err = u.UserRepository.Save(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
