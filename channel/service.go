package channel

type Service interface {
	// FetchAll gets a list of all active channels
	FetchAll() (Channels, error)
	// FetchAllById gets a list of all active channels under a specified world
	FetchAllbyId(worldId byte) (Channels, error)
	// SetAdultChannel changes the adult flag to true or false for the specified channel id
	SetAdultChannel(id int, flag bool)
}

// channelService provides a rough implementation of ChannelService with a hard-coded datastore
type channelService struct {
}

func NewChannelService() *channelService {
	return &channelService{}
}

// FetchAll gets a list of all active channels
func (r channelService) FetchAll() (Channels, error) {
	return Channels{
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
func (s channelService) FetchAllbyId(worldId byte) (res Channels, err error) {
	channels, err := s.FetchAll()
	for _, channel := range channels {
		if channel.WorldId == worldId {
			res = append(res, channel)
		}
	}

	return
}

// SetAdultChannel changes the adult flag to true or false for the specified channel id
func (s *channelService) SetAdultChannel(id int, flag bool) {

}
