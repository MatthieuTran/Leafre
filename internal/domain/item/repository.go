package item

import (
	"context"
)

// ItemRepository provides an interface for accessing the data-layer
type ItemRepository interface {
	Add(ctx context.Context, item Item) error
	GetAllByCharacterID(ctx context.Context, charID uint32) ([]Item, error)
	Update(ctx context.Context, item Item) error
	Destroy(ctx context.Context, id uint32) (err error)
}
