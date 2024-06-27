package usecase

import (
	"context"
	interfaces "github.com/beriloqueiroz/desafio-back/cache_sync_service/internal/usecase/interfaces"
	"log/slog"
)

type SyncUseCase struct {
	LocationRepository interfaces.LocationRepository
	CacheRepository    interfaces.CacheRepository
	MessageGateway     interfaces.MessageGateway
}

func NewSyncUseCase(
	locationRepository interfaces.LocationRepository,
	cacheRepository interfaces.CacheRepository,
	messageGateway interfaces.MessageGateway) *SyncUseCase {
	return &SyncUseCase{
		LocationRepository: locationRepository,
		CacheRepository:    cacheRepository,
		MessageGateway:     messageGateway,
	}
}

func (u *SyncUseCase) Execute(ctx context.Context) {
	// captura locations
	page := 1
	size := 500
	for {
		locations, err := u.LocationRepository.ListUniqueLocations(ctx, page, size)
		if err != nil {
			slog.Error(err.Error())
			break
		}
		for _, location := range locations {
			message, err := u.MessageGateway.MessageByLocation(ctx, location.City, location.State)
			if err != nil {
				slog.Error(err.Error())
				continue
			}
			err = u.CacheRepository.Save(ctx, location.String(), message)
			if err != nil {
				slog.Error(err.Error())
			}
		}
		if len(locations) < size {
			break
		}
		page++
	}
	// captura mensagens
	// popula cache
}
