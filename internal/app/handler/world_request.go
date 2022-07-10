package handler

import (
	"context"
	"io"
	"log"

	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/internal/domain"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeWorldRequest uint16 = 0xB

type HandlerWorldRequest struct {
	worldChannelService domain.WorldChannelService
}

func NewHandlerWorldRequest(worldChannelService domain.WorldChannelService) HandlerWorldRequest {
	return HandlerWorldRequest{
		worldChannelService: worldChannelService,
	}
}

func (h *HandlerWorldRequest) Handle(w io.Writer, p packet.Packet) {
	// Fetch a list of all the worlds
	worlds, err := h.worldChannelService.GetAllWorlds(context.Background())
	if err != nil {
		log.Println("Cannot fetch worlds:", err)
		return
	}

	// Send information regarding the world and its channels
	for _, worldChannel := range worlds {
		writer.WriteWorldInformation(w, worldChannel.World, worldChannel.Channels)
	}

	// Send the signal that there are no more worlds to be added
	writer.WriteWorldInformationDone(w)

	// Send the user's last connected world
	send := writer.SendLatestConnectedWorld{
		LatestConnectedWorld: 0,
	}
	writer.WriteLatestConnectedWorld(w, send)
}

func (h *HandlerWorldRequest) String() string {
	return "WorldRequest"
}
