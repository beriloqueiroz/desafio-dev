package implements

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/beriloqueiroz/desafio-back/core/internal/entity"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log/slog"
)

type WebKafkaRepository struct {
	Producer     *kafka.Producer
	Topic        string
	DeliveryChan chan kafka.Event
}

func NewWebKafkaRepository(producer *kafka.Producer, topic string) *WebKafkaRepository {
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
		Producer:     producer,
		Topic:        topic,
		DeliveryChan: deliveryChan,
	}
}

func (k *WebKafkaRepository) Send(ctx context.Context, notification *entity.Notification) error {
	value, err := json.Marshal(notification)
	if err != nil {
		slog.Error("Erro ao tentar serializar para kafka", "error", err)
		return err
	}
	err = k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.Topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, k.DeliveryChan)
	if err != nil {
		slog.Error("Erro ao tentar enviar para kafka", "error", err)
		return err
	}
	return nil
}
