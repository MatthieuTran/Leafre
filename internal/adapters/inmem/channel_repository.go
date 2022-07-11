package inmem

import (
	"context"
	"strconv"

	"github.com/matthieutran/leafre-login/internal/domain/channel"
)

// ChannelRepository implements `channel.ChannelRepository` with an in-memory map
type ChannelRepository struct {
	channelDict map[string]channel.Channel
}

func NewChannelRepository() channel.ChannelRepository {
	// Create map of channel id -> channel
	channelDict := make(map[string]channel.Channel)
	r := &ChannelRepository{channelDict: channelDict}

	// Store initial data
	dummy_channel := channel.Channel{
		UserNo:       0,
		WorldID:      0,
		ChannelID:    0,
		AdultChannel: 0,
	}
	dummy_channel2 := channel.Channel{
		UserNo:       0,
		WorldID:      0,
		ChannelID:    1,
		AdultChannel: 0,
	}
	r.Add(context.Background(), dummy_channel)
	r.Add(context.Background(), dummy_channel2)

	return &ChannelRepository{channelDict: channelDict}
}

func (r ChannelRepository) Add(ctx context.Context, ch channel.Channel) error {
	ch.ID = strconv.Itoa(len(r.channelDict))
	r.channelDict[ch.ID] = ch
	return nil
}

func (r ChannelRepository) GetAll(ctx context.Context) (channel.Channels, error) {
	res := make(channel.Channels, len(r.channelDict))
	for _, ch := range r.channelDict {
		res = append(res, ch)
	}

	return res, nil
}

func (r ChannelRepository) GetAllByWorldID(ctx context.Context, worldID byte) (channel.Channels, error) {
	var res channel.Channels
	for _, ch := range r.channelDict {
		if ch.WorldID == worldID {
			res = append(res, ch)
		}
	}

	return res, nil
}

func (r ChannelRepository) Update(ctx context.Context, channel channel.Channel) error {
	return nil
}

func (r ChannelRepository) Destroy(ctx context.Context, id string) (err error) {
	delete(r.channelDict, id)
	return nil
}
