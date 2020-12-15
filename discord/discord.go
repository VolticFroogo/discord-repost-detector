package discord

import (
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	Discord *discordgo.Session
)

func Init() {
	var err error
	Discord, err = discordgo.New("Bot " + model.DiscordToken)
	if err != nil {
		panic(err)
	}

	Discord.AddHandler(messageCreate)

	Discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = Discord.Open()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to Discord.")
}

func Close() {
	err := Discord.Close()
	if err != nil {
		panic(err)
	}

	log.Println("Disconnected from Discord.")
}
