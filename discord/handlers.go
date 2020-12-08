package discord

import (
	"github.com/VolticFroogo/discord-repost-detector/match"
	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.Attachments) == 0 {
		return
	}

	for _, attachment := range m.Attachments {
		match.Check(s, m, attachment)
	}
}
