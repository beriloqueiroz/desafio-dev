package usecase

import (
	"context"
	"errors"
	"github.com/beriloqueiroz/desafio-back/core/internal/entity"
	"github.com/beriloqueiroz/desafio-back/core/internal/usecase/interfaces"
	"github.com/google/uuid"
	"time"
)

type CreateScheduleNotificationUseCase struct {
	ScheduleRepository interfaces.ScheduleNotificationRepository
}

func (u *CreateScheduleNotificationUseCase) Execute(ctx context.Context, startTime time.Time) error {
	if startTime.Before(time.Now()) {
		return errors.New("start time is earlier than current time")
	}
	scheduleNotification, err := entity.NewScheduleNotification(uuid.NewString(), startTime, entity.Pending)
	if err != nil {
		return err
	}
	err = u.ScheduleRepository.Save(ctx, scheduleNotification)
	if err != nil {
		return err
	}
	return nil
}
