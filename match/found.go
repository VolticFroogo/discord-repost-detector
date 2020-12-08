package match

import (
	"fmt"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
	"github.com/xeonx/timeago"
	"log"
	"strconv"
)

func found(s *discordgo.Session, m *discordgo.MessageCreate, matches []model.Image) {
	embed := model.DefaultEmbed(m.Author)
	embed.Title = "Repost Detected"

	var pluralS string
	if len(matches) == 1 {
		pluralS = ""
	} else {
		pluralS = "s"
	}

	embed.Description = fmt.Sprintf("I've seen this image in this server %d time%s.", len(matches), pluralS)

	firstUser, err := s.User(strconv.FormatUint(matches[0].User, 10))
	if err != nil {
		log.Printf("Could not find user: %s", err)

		firstUser = &discordgo.User{
			Username: "unknown user",
		}
	}

	embed.Fields = []*discordgo.MessageEmbedField{
		generateField(matches[0], firstUser.Username, "I first saw this image"),
	}

	if len(matches) != 1 {
		lastUser, err := s.User(strconv.FormatUint(matches[len(matches)-1].User, 10))
		if err != nil {
			log.Printf("Could not find user: %s", err)

			lastUser = &discordgo.User{
				Username: "unknown user",
			}
		}

		embed.Fields = append(embed.Fields, generateField(matches[len(matches)-1], lastUser.Username, "I last saw this image"))
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		log.Printf("Error sending embed message: %s", err)
	}
}

func generateField(match model.Image, username, name string) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:  name,
		Value: fmt.Sprintf("[%s posted by %s](https://discordapp.com/channels/%d/%d/%d)", timeago.English.Format(match.DateTime.Time()), username, match.Guild, match.Channel, match.Message),
	}
}
