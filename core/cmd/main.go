package main

import (
	"context"
	"database/sql"
	"github.com/beriloqueiroz/desafio-back/cache_sync_service/pkg/factories"
	"github.com/beriloqueiroz/desafio-back/core/configs"
	"github.com/beriloqueiroz/desafio-back/core/internal/infra/implements"
	"github.com/beriloqueiroz/desafio-back/core/internal/infra/web"
	"github.com/beriloqueiroz/desafio-back/core/internal/usecase"
	"github.com/beriloqueiroz/desafio-back/core/internal/usecase/interfaces"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"os/signal"
	"time"
)

//	@title			Swagger Desafio Meli API
//	@version		1.0
//	@description	This is a notification server .
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	berilo.queiroz@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

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

	kafkaNotificationQueueProduce, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configs.KAFKAUrl,
		"client.id":         "desafio-de-back",
		"acks":              "all",
	})
	if err != nil {
		panic(err)
	}
	defer kafkaNotificationQueueProduce.Close()

	// repositories
	userRepository := implements.PostgresUserRepository{
		Db: db,
	}
	scheduleRepository := implements.PostgresScheduleRepository{
		Db: db,
	}

	messageGateway := factories.NewGetMessageUseCase()
	webKafkaRepository := implements.NewWebKafkaRepository(kafkaNotificationQueueProduce, configs.KAFKATopic)
	notificationQueueRepositories := []interfaces.NotificationQueueRepository{
		webKafkaRepository,
	}

	// useCases
	insertUserUseCase := usecase.InsertUserUseCase{
		UserRepository: &userRepository,
	}
	activateUserRepository := usecase.ActivateUserUseCase{
		UserRepository: &userRepository,
	}
	deactivateUserRepository := usecase.DeactivateUserUseCase{
		UserRepository: &userRepository,
	}

	createScheduleNotificationUseCase := usecase.CreateScheduleNotificationUseCase{
		ScheduleRepository: &scheduleRepository,
	}
	deleteScheduleNotificationUseCase := usecase.DeleteScheduleNotificationUseCase{
		ScheduleRepository: &scheduleRepository,
	}
	syncSchedulesUseCase := usecase.SyncSchedulesNotificationUseCase{
		ScheduleRepository: &scheduleRepository,
		UserRepository:     &userRepository,
		MessageGateway:     messageGateway,
		NotificationQueues: notificationQueueRepositories,
	}

	// webserver and routes
	port := ":8080"
	webserver := web.NewWebServer(port)

	userRoutes := web.NewUserRoutes(insertUserUseCase, activateUserRepository, deactivateUserRepository)
	webserver.AddRoute("POST /user", userRoutes.InsertUserHandler)
	webserver.AddRoute("PUT /user/{id}/activate", userRoutes.ActivateUserHandler)
	webserver.AddRoute("PUT /user/{id}/deactivate", userRoutes.DeactivateUserHandler)

	scheduleRoutes := web.NewSchedulerRoutes(createScheduleNotificationUseCase, deleteScheduleNotificationUseCase)
	webserver.AddRoute("POST /schedule", scheduleRoutes.CreateScheduleNotificationHandler)
	webserver.AddRoute("DELETE /schedule/{id}", scheduleRoutes.DeleteScheduleNotificationHandler)

	// start server
	srvErr := make(chan error, 1)
	go func() {
		slog.Info("Starting web server", "on port", port)
		srvErr <- webserver.Start()
	}()

	// jobs
	go func() {
		for {
			slog.Info("Starting sync schedules")
			err := syncSchedulesUseCase.Execute(context.Background())
			if err != nil {

				slog.Error(err.Error())
			}
			time.Sleep(time.Second)
		}
	}()

	// Wait for interruption.
	select {
	case <-sigCh:
		slog.Warn("Shutting down gracefully, CTRL+C pressed...")
	case <-initCtx.Done():
		slog.Warn("Shutting down due to other reason...")
	}
}
