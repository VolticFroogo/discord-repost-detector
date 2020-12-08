package model

import "github.com/bwmarrin/discordgo"

func DefaultEmbed(author *discordgo.User) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 1029388,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Discord Repost Detector by Froogo",
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    author.Username,
			IconURL: author.AvatarURL(""),
		},
	}
}
