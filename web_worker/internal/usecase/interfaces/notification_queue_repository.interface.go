package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-back/web_worker/internal/entity"
)

type Action func(ctx context.Context, notifications []entity.Notification) error

type NotificationQueueRepository interface {
	Read(ctx context.Context, action Action) error
}
