package handler

import (
	"io"
	"log"

	"github.com/matthieutran/duey"
	login "github.com/matthieutran/leafre-login"
	"github.com/matthieutran/packet"
)

const OpCodeWorldRequest uint16 = 0xB

type HandlerWorldRequest struct {
	worldRepository login.WorldRepository
}

func NewHandlerWorldRequest(worldRepository login.WorldRepository) HandlerWorldRequest {
	return HandlerWorldRequest{
		worldRepository: worldRepository,
	}
}

func (h *HandlerWorldRequest) Handle(w io.Writer, es *duey.EventStreamer, p packet.Packet) {
	worlds, err := h.worldRepository.FetchAll()
	if err != nil {
		log.Println("Cannot fetch worlds!")
		return
	}

	for world := range worlds {
		log.Println(world)
	}
}

func (h *HandlerWorldRequest) String() string {
	return "WorldRequest"
}
