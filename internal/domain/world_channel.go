package domain

import (
	"context"
	"log"

	"github.com/matthieutran/leafre-login/internal/domain/channel"
	"github.com/matthieutran/leafre-login/internal/domain/world"
)

// WorldChannel contains a World and its Channel
type WorldChannel struct {
	world.World
	channel.Channels
}

type WorldChannels []WorldChannel

type WorldChannelService interface {
	GetAllWorlds(context.Context) (WorldChannels, error)
	SetAdultChannel(ctx context.Context, channelID string, flag byte)
}

func NewWorldChannelService(worldRepo world.WorldRepository, channelRepo channel.ChannelRepository) WorldChannelService {
	return defaultWorldChannelService{worldRepo: worldRepo, channelRepo: channelRepo}
}

type defaultWorldChannelService struct {
	worldRepo   world.WorldRepository
	channelRepo channel.ChannelRepository
}

func (s defaultWorldChannelService) GetAllWorlds(ctx context.Context) (res WorldChannels, err error) {
	worlds, err := s.worldRepo.GetAll(ctx)
	if err != nil {
		return
	}

	for _, world := range worlds {
		channels, err := s.channelRepo.GetAllByWorldID(ctx, world.ID)
		if err != nil {
			log.Printf("Could not get channels for world (%d): %s", world.ID, err)
		}

		wc := WorldChannel{
			World:    world,
			Channels: channels,
		}

		res = append(res, wc)
	}

	return
}

func (s defaultWorldChannelService) SetAdultChannel(ctx context.Context, channelID string, flag byte) {

}
