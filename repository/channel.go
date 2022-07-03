package repository

import login "github.com/matthieutran/leafre-login"

// ChannelRepository provides the business methods to the channel object
type ChannelRepository struct {
}

func NewChannelRepostory() ChannelRepository {
	return ChannelRepository{}
}

// FetchAll gets a list of all active channels
func (r ChannelRepository) FetchAll() (login.Channels, error) {
	return login.Channels{
		{ // Scania, Ch 1
			Id:           "0",
			UserNo:       0,
			WorldId:      0,
			ChannelId:    0,
			AdultChannel: 0,
		},
		{ // SomeOtherWorld, Ch 1
			Id:           "1",
			UserNo:       0,
			WorldId:      1, // NOTE different world id ^
			ChannelId:    0,
			AdultChannel: 0,
		},
	}, nil
}

// FetchAllById gets a list of all active channels under a specified world
func (r ChannelRepository) FetchAllbyId(worldId byte) (res login.Channels, err error) {
	channels, err := r.FetchAll()
	for _, channel := range channels {
		if channel.WorldId == worldId {
			res = append(res, channel)
		}
	}

	return
}

// SetAdultChannel changes the adult flag to true or false for the specified channel id
func (r ChannelRepository) SetAdultChannel(id int, flag bool) {

}
