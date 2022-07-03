package repository

import login "github.com/matthieutran/leafre-login"

// ChannelRepository provides the business methods to the channel object
type ChannelRepository struct {
}

func NewChannelRepostory() ChannelRepository {
	return ChannelRepository{}
}

// FetchAll gets a list of all active channels
func FetchAll() (login.Channels, error) {
	return login.Channels{
		{
			Id:           "0",
			UserNo:       0,
			WorldId:      0,
			ChannelId:    0,
			AdultChannel: false,
		},
	}, nil
}

// SetAdultChannel changes the adult flag to true or false for the specified channel id
func SetAdultChannel(id int, flag bool) {

}
