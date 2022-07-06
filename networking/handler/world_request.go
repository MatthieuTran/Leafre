package handler

import (
	"io"
	"log"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/channel"
	"github.com/matthieutran/leafre-login/networking/writer"
	"github.com/matthieutran/leafre-login/world"
	"github.com/matthieutran/packet"
)

const OpCodeWorldRequest uint16 = 0xB

type HandlerWorldRequest struct {
	worldService   world.WorldService
	channelService channel.ChannelService
}

func NewHandlerWorldRequest(worldService world.WorldService, channelService channel.ChannelService) HandlerWorldRequest {
	return HandlerWorldRequest{
		worldService:   worldService,
		channelService: channelService,
	}
}

func (h *HandlerWorldRequest) Handle(w io.Writer, es *duey.EventStreamer, p packet.Packet) {
	// Fetch a list of all the worlds
	worlds, err := h.worldService.FetchAll()
	if err != nil {
		log.Println("Cannot fetch worlds!")
		return
	}

	// Start sending world information
	for _, world := range worlds {
		channels, err := h.channelService.FetchAllbyId(world.Id)
		if err != nil {
			log.Printf("Could not get channel information (World ID: %d)", world.Id)
		}

		// Send information regarding the world and its channels
		writer.WriteWorldInformation(w, world, channels)
	}

	// Send the signal that there are no more worlds to be added
	writer.WriteWorldInformationDone(w)

	// Send the user's last connected world
	latestConnectedWorld := 0
	writer.WriteLatestConnectedWorld(w, latestConnectedWorld)
}

func (h *HandlerWorldRequest) String() string {
	return "WorldRequest"
}
