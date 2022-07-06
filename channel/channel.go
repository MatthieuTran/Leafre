package channel

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
	AdultChannel byte `json:"adult_channel"`
}

type Channels []Channel
