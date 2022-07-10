package inmem

import (
	"context"
	"log"

	"github.com/matthieutran/leafre-login/internal/domain/world"
)

// WorldRepository implements `world.WorldRepository` with an in-memory map
type WorldRepository struct {
	worldDict map[byte]world.World
}

func NewWorldRepository() world.WorldRepository {
	// Create map of world id -> world
	worldDict := make(map[byte]world.World)
	r := &WorldRepository{worldDict: worldDict}

	// Store initial data
	dummy_world := world.World{
		Name:      "Scania",
		EventDesc: "Leafre",
	}
	r.Add(context.Background(), dummy_world)
	log.Println("Added world", dummy_world)

	return &WorldRepository{worldDict: worldDict}
}

func (r WorldRepository) Add(ctx context.Context, w world.World) error {
	w.ID = byte(len(r.worldDict))
	r.worldDict[w.ID] = w
	return nil
}

func (r WorldRepository) GetAll(ctx context.Context) (worlds world.Worlds, err error) {
	log.Println("PRE:", r.worldDict)
	for _, ch := range r.worldDict {
		worlds = append(worlds, ch)
		log.Println("TEST", worlds)
	}

	return
}

func (r WorldRepository) GetByID(ctx context.Context, id byte) (world.World, error) {
	for _, w := range r.worldDict {
		if w.ID == id {
			return w, nil
		}
	}

	return world.World{}, world.ErrDoesNotExist
}
