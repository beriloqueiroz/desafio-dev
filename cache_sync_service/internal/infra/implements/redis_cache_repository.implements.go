package implements

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCacheRepository struct {
	Client *redis.Client
}

func NewRedisCacheRepository(client *redis.Client) *RedisCacheRepository {
	return &RedisCacheRepository{
		Client: client,
	}
}

func (r *RedisCacheRepository) Find(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisCacheRepository) Save(ctx context.Context, key string, value string) error {
	now := time.Now()
	yyyy, mm, dd := now.Date()
	tomorrow := time.Date(yyyy, mm, dd+1, 0, 0, 0, 0, now.Location()) // todo a hora pode ser variável de ambiente
	timeToExpire := tomorrow.Sub(now)
	err := r.Client.Set(ctx, key, value, timeToExpire).Err() // expiração pode ser variável de ambiente
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCacheRepository) SaveAll(ctx context.Context, values map[string]string) error {
	for k, v := range values {
		err := r.Client.HSet(ctx, time.Now().String(), k, v).Err()
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (r *RedisCacheRepository) Delete(ctx context.Context, key string) error {
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
