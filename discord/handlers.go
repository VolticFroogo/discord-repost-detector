package discord

import (
	"github.com/VolticFroogo/discord-repost-detector/command"
	"github.com/VolticFroogo/discord-repost-detector/db"
	"github.com/VolticFroogo/discord-repost-detector/match"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
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

	channel, err := strconv.ParseUint(m.ChannelID, 0, 64)
	if err != nil {
		log.Printf("Error parsing guild ID: %s", err)
		return
	}

	res := db.Channels.FindOne(db.DefaultContext(), bson.M{
		"_id": channel,
	})
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		log.Printf("Error finding channel: %s", res.Err())
		return
	}

	if res.Err() == nil {
		for _, attachment := range m.Attachments {
			match.Check(s, m, attachment)
		}
	}
}
