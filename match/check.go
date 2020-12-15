package match

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Check(s *discordgo.Session, m *discordgo.MessageCreate, attachment *discordgo.MessageAttachment) {
	// We only want images, so skip over any other files.
	if attachment.Width == 0 || attachment.Height == 0 {
		return
	}

	// Download image.
	res, err := http.Get(attachment.ProxyURL)
	if err != nil {
		panic(err)
	}

	// Read image into bytes.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Hash the bytes.
	hash := md5.Sum(body)
	hashBase64 := base64.StdEncoding.EncodeToString(hash[:])

	guild, err := strconv.ParseUint(m.GuildID, 0, 64)
	if err != nil {
		panic(err)
	}

	channel, err := strconv.ParseUint(m.ChannelID, 0, 64)
	if err != nil {
		panic(err)
	}

	message, err := strconv.ParseUint(m.ID, 0, 64)
	if err != nil {
		panic(err)
	}

	user, err := strconv.ParseUint(m.Author.ID, 0, 64)
	if err != nil {
		panic(err)
	}

	image := model.Image{
		Hash:    hashBase64,
		Guild:   guild,
		Channel: channel,
		Message: message,
		User:    user,
	}

	matches := image.FindMatch()

	if len(matches) != 0 {
		found(s, m, matches)
	}

	image.Insert()
}
