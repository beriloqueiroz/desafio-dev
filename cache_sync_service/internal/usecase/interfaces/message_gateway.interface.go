package interfaces

import "context"

type MessageGateway interface {
	MessageByLocation(ctx context.Context, city string, state string) (string, error)
}
