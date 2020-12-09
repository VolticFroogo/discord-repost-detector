package main

import (
	"github.com/VolticFroogo/discord-repost-detector/db"
	"github.com/VolticFroogo/discord-repost-detector/discord"
	"github.com/VolticFroogo/discord-repost-detector/status"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Connect to MongoDB.
	err := db.Init()
	if err != nil {
		log.Fatalf("Error initialising DB: %s", err)
	}

	// Connect to Discord.
	err = discord.Init()
	if err != nil {
		log.Fatalf("Error initialising Discord: %s", err)
	}

	// Start the status updater.
	statusQuitChan, statusFinished := status.Start()

	// Wait until we receive an interrupt signal.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Quit out of the status updater.
	statusQuitChan <- true
	<-statusFinished

	// Disconnect from Discord.
	err = discord.Close()
	if err != nil {
		log.Fatalf("Error closing Discord: %s", err)
	}

	// Disconnect from MongoDB.
	err = db.Close()
	if err != nil {
		log.Fatalf("Error closing DB: %s", err)
	}
}
