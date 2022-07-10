package channel

type Channel struct {
	ID           string // Channel ID
	UserNo       uint32 // Number of users currently on the channel (to measure channel popularity)
	WorldID      byte   // World ID
	ChannelID    byte   // Channel ID *in the world*
	AdultChannel byte   // Adult Channel (18+)
}

type Channels []Channel
