package main

import (
	"log"
	"os"
	"sync"

	"github.com/matthieutran/leafre-login/messaging"
	"github.com/matthieutran/leafre-login/server"
)

func main() {
	var wg sync.WaitGroup

	log.Println("Leafre - Login Server")

	es := messaging.Init(os.Getenv("NATS_URI"))

	// Create socket
	s := server.BuildServer(wg, es)
	wg.Add(1)
	s.Start(wg)

	// Block until all goroutines are done
	wg.Wait()
}
