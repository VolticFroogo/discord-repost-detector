package command

import (
	"fmt"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
)

func hasAdministrator(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (admin bool) {
	// NOTE: This seems to be very inefficient as it requires two additional API calls just to check the author's permissions.
	// I couldn't find a better way to solve this using discordgo, so this will have to work for now.
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		panic(err)
	}

	if guild.OwnerID == m.Author.ID {
		admin = true
		return
	}

	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		panic(err)
	}

	j := 0
	for i := 0; i < len(guild.Roles); i++ {
		if guild.Roles[i].ID != member.Roles[j] {
			continue
		}

		if guild.Roles[i].Permissions&discordgo.PermissionAdministrator != 0 {
			admin = true
			return
		}

		j++
	}

	badPermissions(s, m, args)
	return
}

func badPermissions(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	embed := model.DefaultEmbed(m.Author)
	embed.Title = "Unauthorised"
	embed.Description = fmt.Sprintf("You must have the Administrator permission to execute the %s command.", args[1])

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		panic(err)
	}
}
