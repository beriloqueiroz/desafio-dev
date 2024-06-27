package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-back/core/internal/usecase/interfaces"
)

type DeactivateUserUseCase struct {
	UserRepository interfaces.UserRepository
}

func (u *DeactivateUserUseCase) Execute(ctx context.Context, ID string) error {
	user, err := u.UserRepository.Find(ctx, ID)
	if err != nil {
		return err
	}
	user.Deactivate()
	err = u.UserRepository.Save(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
