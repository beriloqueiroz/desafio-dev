package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-back/core/internal/entity"
)

type UserRepository interface {
	Find(ctx context.Context, id string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
	ListActives(ctx context.Context, page, size int) ([]entity.User, error)
}
