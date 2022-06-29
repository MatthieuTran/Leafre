package main

import (
	"log"
	"os"
	"sync"

	"github.com/matthieutran/leafre-login/internal/login/messaging"
	"github.com/matthieutran/leafre-login/internal/login/net"
)

func main() {
	var wg sync.WaitGroup

	log.Println("Leafre - Login Server")

	s := messaging.Init(os.Getenv("NATS_URI"))

	// Create socket
	server := net.BuildServer(wg, s)
	wg.Add(1)
	server.Start(wg)

	// Block until all goroutines are done
	wg.Wait()
}
