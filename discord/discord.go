package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

var (
	discord *discordgo.Session
)

func Init() (err error) {
	token := os.Getenv("DISCORD_TOKEN")
	discord, err = discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("creating discord client: %s", err)
	}

	discord.AddHandler(messageCreate)

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = discord.Open()
	if err != nil {
		return fmt.Errorf("connecting to Discord: %s", err)
	}

	log.Println("Connected to Discord.")

	return
}

func Close() (err error) {
	err = discord.Close()
	if err != nil {
		return fmt.Errorf("disconnecting: %s", err)
	}

	return
}
