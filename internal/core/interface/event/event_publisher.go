package event

import "context"

type EventPublisher interface {
	Publish(ctx context.Context, event interface{}) error
}
