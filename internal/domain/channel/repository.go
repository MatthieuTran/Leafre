package channel

import (
	"context"
)

// ChannelRepository provides an interface for accessing the data-layer
type ChannelRepository interface {
	Add(ctx context.Context, world Channel) error
	GetAll(ctx context.Context) (Channels, error)
	GetAllByWorldID(ctx context.Context, worldID byte) (Channels, error)
	Update(ctx context.Context, world Channel) error
	Destroy(ctx context.Context, id string) (err error)
}
