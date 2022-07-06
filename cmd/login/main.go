package main

import (
	"log"
	"os"
	"sync"

	"github.com/matthieutran/leafre-login/messaging"
	"github.com/matthieutran/leafre-login/networking"
)

func main() {
	var wg sync.WaitGroup

	log.Println("Leafre - Login Server")

	es := messaging.Init(os.Getenv("NATS_URI"))

	// Create socket
	wg.Add(1)
	s := networking.BuildServer(&wg, es)
	s.Start(&wg)

	// Block until all goroutines are done
	wg.Wait()
}
