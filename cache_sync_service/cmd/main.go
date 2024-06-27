package main

import (
	"context"
	"database/sql"
	"github.com/beriloqueiroz/desafio-back/cache_sync_service/configs"
	implements2 "github.com/beriloqueiroz/desafio-back/cache_sync_service/internal/infra/implements"
	"github.com/beriloqueiroz/desafio-back/cache_sync_service/internal/usecase"
	_ "github.com/lib/pq"
	redis "github.com/redis/go-redis/v9"
	"github.com/robfig/cron"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	// graceful exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	initCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// config logs
	// aqui no lugar do Stdout poderia ser um db ou elasticsearch por exemplo
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// load environment configs
	configs, err := configs.LoadConfig([]string{"."})
	if err != nil {
		panic(err)
	}

	// dbs
	db, err := sql.Open(configs.DBDriver, configs.DBUri)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     configs.CacheUri,
		Password: configs.CachePass,
		DB:       0,
	})

	// repositories e gateways
	redisCacheRepository := implements2.NewRedisCacheRepository(clientRedis)
	locationRepository := &implements2.PostgresLocationRepository{
		Db: db,
	}
	messageGateway := implements2.NewCptecMessageGateway()

	syncUseCase := usecase.NewSyncUseCase(locationRepository, redisCacheRepository, messageGateway)

	c := cron.New()
	err = c.AddFunc("0 0 0 * * *", func() { // todo a hora pode ser variável de ambiente, a mesma do timeToExpire
		slog.Info("Starting sync cache")
		syncUseCase.Execute(context.Background())
		slog.Info("End sync cache")
	})
	if err != nil {
		panic(err)
	}
	c.Start()

	// Wait for interruption.
	select {
	case <-sigCh:
		slog.Warn("Shutting down gracefully, CTRL+C pressed...")
	case <-initCtx.Done():
		slog.Warn("Shutting down due to other reason...")
	}
}
