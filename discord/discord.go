package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

var (
	Discord *discordgo.Session
)

func Init() (err error) {
	token := os.Getenv("DISCORD_TOKEN")
	Discord, err = discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("creating Discord client: %s", err)
	}

	Discord.AddHandler(messageCreate)

	Discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = Discord.Open()
	if err != nil {
		return fmt.Errorf("connecting to Discord: %s", err)
	}

	log.Println("Connected to Discord.")

	return
}

func Close() (err error) {
	err = Discord.Close()
	if err != nil {
		return fmt.Errorf("disconnecting: %s", err)
	}

	log.Println("Disconnected from Discord.")
	return
}
