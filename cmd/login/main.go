package main

import (
	"log"
	"sync"

	"github.com/matthieutran/leafre-login/internal/login/net"
)

func main() {
	var wg sync.WaitGroup

	log.Println("Leafre - Login Server")

	// Create socket
	server := net.BuildServer(wg)
	wg.Add(1)
	server.Start(wg)

	// Block until all goroutines are done
	wg.Wait()
}
