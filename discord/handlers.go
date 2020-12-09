package discord

import (
	"github.com/VolticFroogo/discord-repost-detector/command"
	"github.com/VolticFroogo/discord-repost-detector/match"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, model.Ping) {
		err := command.Handle(s, m)
		if err != nil {
			log.Printf("Error handling command: %s", err)
		}

		return
	}

	for _, attachment := range m.Attachments {
		match.Check(s, m, attachment)
	}
}
