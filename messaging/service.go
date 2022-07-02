package messaging

import (
	"log"

	"github.com/matthieutran/duey"
)

func Init(uri string) *duey.EventStreamer {
	s, err := duey.Init(uri)
	if err != nil {
		log.Fatal("Could not connect to messaging system:", err)
	}

	subscribers := []func() (string, duey.Handler){
		// subs
	}

	for _, subscriber := range subscribers {
		s.Subscribe(subscriber())
	}

	return s
}
