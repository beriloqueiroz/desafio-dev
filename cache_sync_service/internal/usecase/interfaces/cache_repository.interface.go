package interfaces

import "context"

type CacheRepository interface {
	Find(ctx context.Context, key string) (string, error)
	Save(ctx context.Context, key string, value string) error
	SaveAll(ctx context.Context, values map[string]string) error
	Delete(ctx context.Context, key string) error
}
