package main

import (
	"github.com/VolticFroogo/discord-repost-detector/command"
	"github.com/VolticFroogo/discord-repost-detector/db"
	"github.com/VolticFroogo/discord-repost-detector/discord"
	"github.com/VolticFroogo/discord-repost-detector/status"
	"github.com/getsentry/sentry-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		ServerName: os.Getenv("SERVER_NAME"),
		Debug:      true,
	})
	if err != nil {
		log.Fatalf("Error initialising Sentry: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(10 * time.Second)
	defer sentry.Recover()

	// Connect to MongoDB.
	db.Init()

	// Connect to Discord.
	discord.Init()

	// Start the status updater.
	statusQuitChan, statusFinished := status.Start()

	// Register our commands.
	command.RegisterCommands()

	// Wait until we receive an interrupt signal.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Quit out of the status updater.
	statusQuitChan <- true
	<-statusFinished

	// Disconnect from Discord.
	discord.Close()

	// Disconnect from MongoDB.
	db.Close()
}
