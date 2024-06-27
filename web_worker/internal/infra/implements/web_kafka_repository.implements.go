package implements

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/beriloqueiroz/desafio-back/web_worker/internal/entity"
	"github.com/beriloqueiroz/desafio-back/web_worker/internal/usecase/interfaces"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log/slog"
	"time"
)

type WebKafkaRepository struct {
	Consumer     *kafka.Consumer
	Topic        string
	DeliveryChan chan kafka.Event
}

func NewWebKafkaRepository(consumer *kafka.Consumer, topic string) *WebKafkaRepository {
	deliveryChan := make(chan kafka.Event, 1000)
	go func() {
		for e := range deliveryChan {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					slog.Error(fmt.Sprintf("Failed to deliver message: %v\n", ev.TopicPartition))
				} else {
					slog.Info(fmt.Sprintf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset))
				}
			}
		}
	}()
	return &WebKafkaRepository{
		Consumer:     consumer,
		Topic:        topic,
		DeliveryChan: deliveryChan,
	}
}

func (k *WebKafkaRepository) Read(ctx context.Context, action interfaces.Action) error {
	run := true
	//defer k.Consumer.Close()
	var notifications []entity.Notification
	for run {
		ev := k.Consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			var notification entity.Notification
			err := json.Unmarshal(e.Value, &notification)
			if err != nil {
				slog.Error(fmt.Sprintf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value)))
				break
			}
			notifications = append(notifications, notification)
		case kafka.PartitionEOF:
			msg := fmt.Sprintf("%% Reached %v\n", e)
			slog.Info(msg)
		case kafka.Error:
			run = false
		default:
			if len(notifications) > 0 {
				err := action(ctx, notifications)
				if err != nil {
					slog.Error(err.Error())
					time.Sleep(time.Second)
					continue
				}
				_, err = k.Consumer.Commit()
				if err != nil {
					slog.Error(err.Error())
					continue
				}
				slog.Info("Send ok")
				notifications = notifications[:0]
			}
		}
	}
	return nil
}
