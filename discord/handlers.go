package discord

import (
	"github.com/VolticFroogo/discord-repost-detector/command"
	"github.com/VolticFroogo/discord-repost-detector/db"
	"github.com/VolticFroogo/discord-repost-detector/match"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/sentry-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer sentry.Recover()

	if m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, model.Ping) {
		command.Handle(s, m)
		return
	}

	channel, err := strconv.ParseUint(m.ChannelID, 0, 64)
	if err != nil {
		panic(err)
	}

	res := db.Channels.FindOne(db.DefaultContext(), bson.M{
		"_id": channel,
	})
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		panic(err)
	}

	if res.Err() == nil {
		for _, attachment := range m.Attachments {
			match.Check(s, m, attachment)
		}
	}
}
