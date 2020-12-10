package command

import (
	"fmt"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
)

func help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (err error) {
	embed := model.DefaultEmbed(m.Author)

	if len(args) > 2 {
		if command, ok := triggerMap[args[2]]; ok {
			embed.Title = fmt.Sprintf("Help for the %s command", command.Name)
			embed.Description = fmt.Sprintf("%s\n\nFormat: %s %s %s\nExample: %s %s %s", command.Description,
				model.Ping, command.Triggers[0], command.Format,
				model.Ping, command.Triggers[0], command.Example)

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			return
		}
	}

	embed.Title = "Help"
	embed.Description = "Repost Detector is a Discord bot which automatically detects and flags images which have been posted in the server before."
	embed.Fields = []*discordgo.MessageEmbedField{
		{
			Name:  "Get started",
			Value: "The simplest way to get started is to view [the getting started page in our documentation](https://github.com/VolticFroogo/discord-repost-detector).",
		},
		{
			Name:  "Commands",
			Value: fmt.Sprintf("To view a list of my commands, type %s commands", model.Ping),
		},
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return
}

func unknownCommand(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) (err error) {
	embed := model.DefaultEmbed(m.Author)
	embed.Title = "Unknown Command"
	embed.Description = fmt.Sprintf("I didn't recognise that command; try %s help", model.Ping)

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return
}

func list(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) (err error) {
	embed := model.DefaultEmbed(m.Author)
	embed.Title = "Commands"
	embed.Description = "A list of all available commands."

	for i := range commands {
		triggers := ""
		for j := 0; j < len(commands[i].Triggers)-1; j++ {
			triggers += commands[i].Triggers[j] + ", "
		}
		triggers += commands[i].Triggers[len(commands[i].Triggers)-1]

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s (%s)", commands[i].Name, triggers),
			Value: fmt.Sprintf("%s\n\nFormat: %s %s %s\nExample: %s %s %s", commands[i].Description,
				model.Ping, commands[i].Triggers[0], commands[i].Format,
				model.Ping, commands[i].Triggers[0], commands[i].Example),
		})
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return
}

func badArguments(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (err error) {
	embed := model.DefaultEmbed(m.Author)
	embed.Title = "Bad/not enough arguments"
	embed.Description = fmt.Sprintf("For more information on how to correctly use this command, type %s help %s", model.Ping, args[1])

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return
}
