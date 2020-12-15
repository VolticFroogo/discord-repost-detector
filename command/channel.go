package command

import (
	"fmt"
	"github.com/VolticFroogo/discord-repost-detector/db"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
)

type Channel struct {
	ID    uint64 `bson:"_id"`
	Guild uint64 `bson:"guild"`
}

func channel(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 3 {
		badArguments(s, m, args)
		return
	}

	switch args[2] {
	case "add":
		addChannel(s, m)

	case "remove":
		removeChannel(s, m)

	default:
		badArguments(s, m, args)
	}

	return
}

func addChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	dgoChannel, err := s.Channel(m.ChannelID)
	if err != nil {
		panic(err)
	}

	channel, err := strconv.ParseUint(m.ChannelID, 0, 64)
	if err != nil {
		panic(err)
	}

	guild, err := strconv.ParseUint(m.GuildID, 0, 64)
	if err != nil {
		panic(err)
	}

	embed := model.DefaultEmbed(m.Author)

	_, err = db.Channels.InsertOne(db.DefaultContext(), Channel{
		ID:    channel,
		Guild: guild,
	})
	if err != nil {
		if !strings.Contains(err.Error(), "duplicate key") {
			return
		}

		embed.Title = fmt.Sprintf("Channel #%s already added", dgoChannel.Name)
		embed.Description = "This channel was already added."
	} else {
		embed.Title = fmt.Sprintf("Added channel #%s", dgoChannel.Name)
		embed.Description = "Repost Detector will now listen for reposts in this channel."
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		panic(err)
	}
}

func removeChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	dgoChannel, err := s.Channel(m.ChannelID)
	if err != nil {
		panic(err)
	}

	channel, err := strconv.ParseUint(m.ChannelID, 0, 64)
	if err != nil {
		panic(err)
	}

	res, err := db.Channels.DeleteOne(db.DefaultContext(), bson.M{
		"_id": channel,
	})
	if err != nil {
		panic(err)
	}

	embed := model.DefaultEmbed(m.Author)

	if res.DeletedCount == 1 {
		embed.Title = fmt.Sprintf("Removed channel #%s", dgoChannel.Name)
		embed.Description = "Repost Detector will no longer listen for reposts in this channel."
	} else {
		embed.Title = fmt.Sprintf("Channel #%s was not added", dgoChannel.Name)
		embed.Description = "Repost Detector was already not listening for reposts in this channel."
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		panic(err)
	}
}
