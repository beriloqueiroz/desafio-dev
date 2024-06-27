package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-back/core/internal/entity"
	"github.com/beriloqueiroz/desafio-back/core/internal/usecase/interfaces"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type SyncSchedulesNotificationUseCase struct {
	UserRepository     interfaces.UserRepository
	ScheduleRepository interfaces.ScheduleNotificationRepository
	NotificationQueues []interfaces.NotificationQueueRepository
	MessageGateway     interfaces.MessageGateway
}

func (u *SyncSchedulesNotificationUseCase) Execute(ctx context.Context) error {
	// todo buscar primeiro scheduler não executados com data anterior a atual
	scheduler, err := u.ScheduleRepository.FindFirstPendingBeforeDate(ctx, time.Now())
	if scheduler == nil && err == nil {
		return nil
	}
	if err != nil {
		return err
	}
	slog.Info("Iniciando schedule", "scheduler id", scheduler.ID)
	scheduler.MarkProcessing()
	err = u.ScheduleRepository.Save(ctx, scheduler)
	if err != nil {
		return err
	}
	// todo buscar users ativos com paginação
	page := 1
	size := 500
	for {
		users, err := u.UserRepository.ListActives(ctx, page, size)
		if users == nil && err == nil {
			slog.Warn("Sem usuários para schedule", "schedule id", scheduler.ID)
			scheduler.MarkExecuted()
			break
		}
		if err != nil {
			slog.Error("Falha ao listar users", "scheduler id", scheduler.ID, "error", err.Error())
			scheduler.MarkExecutedWithError()
			break
		}
		uniquesLocations := getUniquesLocation(users)
		// todo buscar mensagens com base nas cidades dos usuários
		locationsMapMsg, err := u.MessageGateway.ListByLocations(ctx, uniquesLocations)
		if err != nil {
			slog.Error("Falha ao listar localidades", "scheduler id", scheduler.ID, "error", err.Error())
			scheduler.MarkExecutedWithError()
			break
		}
		for _, user := range users {
			// todo montar notificações enviar notificações para as filas
			notification, err := entity.NewNotification(uuid.NewString(), user, *scheduler, locationsMapMsg[user.Location.String()])
			if err != nil {
				slog.Error("Falha ao listar users", "scheduler id", scheduler.ID, "user id", user.ID, "error", err.Error())
				scheduler.MarkExecutedWithError()
				continue
			}
			for _, queue := range u.NotificationQueues {
				err = queue.Send(ctx, notification) // todo aqui poderia usar go rotine para enviar em paralelo
				if err != nil {
					slog.Error("Falha ao listar users", "scheduler id", scheduler.ID, "user id", user.ID, "queue", queue, "error", err.Error())
					scheduler.MarkExecutedWithError()
				}
			}
		}
		if len(users) < size {
			if scheduler.Status == entity.Processing {
				scheduler.MarkExecuted()
			}
			break
		}
		page++
	}
	// todo marcar schedulers como executed
	err = u.ScheduleRepository.Save(ctx, scheduler)
	if err != nil {
		return err
	}
	return nil
}

func getUniquesLocation(sliceList []entity.User) []entity.Location {
	allKeys := make(map[string]bool)
	var list []entity.Location
	for _, item := range sliceList {
		if _, value := allKeys[item.Location.String()]; !value {
			allKeys[item.Location.String()] = true
			list = append(list, item.Location)
		}
	}
	return list
}
