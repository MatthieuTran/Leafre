package login

type Channel struct {
	// ID of this channel
	Id string `json:"channel_id"`
	// Number of users currently on the channel (to measure channel popularity)
	UserNo int `json:"user_no"`
	// ID of the world associated
	WorldId byte `json:"world_id"`
	// ID of the channel (according to the world)
	ChannelId byte `json:"world_channel_id"`
	// Adult Channel (18+)
	AdultChannel bool `json:"adult_channel"`
}

type Channels []Channel

type ChannelRepository interface {
	// FetchAll gets a list of all active channels
	FetchAll() (Channels, error)
	// FetchAllById gets a list of all active channels under a specified world
	FetchAllbyId(worldId byte) (Channels, error)
	// SetAdultChannel changes the adult flag to true or false for the specified channel id
	SetAdultChannel(id int, flag bool)
}
