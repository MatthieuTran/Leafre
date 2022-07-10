package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/matthieutran/leafre-login/internal/ports"
	"github.com/matthieutran/leafre-login/internal/service"
)

const (
	SERVICE_NAME = "Login Server"
	HOST         = ""
	PORT         = 8484
)

func main() {
	var wg sync.WaitGroup

	log.Println(SERVICE_NAME, "- Leafre")

	// Listen for interrupt system call
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Create context
	ctx, cancel := context.WithCancel(context.Background())

	// Gracefully shutdown on interrupt
	go func() {
		oscall := <-c
		log.Printf("Shutting down %s... (System call: %+v)", SERVICE_NAME, oscall)
		cancel()
	}()

	// Start up server
	wg.Add(1)
	app := service.NewApplication()
	go ports.StartTCPServer(&wg, ctx, app)(HOST, PORT) // Inject WaitGroup and Context and call with options

	wg.Wait()
}
