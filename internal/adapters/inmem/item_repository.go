package inmem

import (
	"context"

	"github.com/matthieutran/leafre-login/internal/domain/item"
)

// ItemRepository implements `item.ItemRepository` with an in-memory map
type ItemRepository struct {
	items map[uint32]item.Item
}

func NewItemRepository() item.ItemRepository {
	// Create map of item id -> item
	items := make(map[uint32]item.Item)
	r := &ItemRepository{items: items}

	return r
}

func (r ItemRepository) Add(ctx context.Context, i item.Item) error {
	i.ID = uint32(len(r.items))
	r.items[i.ID] = i

	return nil
}

func (r ItemRepository) GetAllByCharacterID(ctx context.Context, charID uint32) (inv item.Items, err error) {
	for _, i := range r.items {
		if i.CharacterID == charID {
			inv = append(inv, i)
		}
	}

	return
}

func (r ItemRepository) Update(ctx context.Context, item item.Item) error {
	return nil
}

func (r ItemRepository) Destroy(ctx context.Context, id uint32) (err error) {
	delete(r.items, id)
	return nil
}
