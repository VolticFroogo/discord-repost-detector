package match

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
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
		log.Printf("Error downloading image: %s", err)
		return
	}

	// Read image into bytes.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading body: %s", err)
		return
	}

	// Hash the bytes.
	hash := md5.Sum(body)
	hashBase64 := base64.StdEncoding.EncodeToString(hash[:])

	guild, err := strconv.ParseUint(m.GuildID, 0, 64)
	if err != nil {
		log.Printf("Error parsing guild ID: %s", err)
		return
	}

	channel, err := strconv.ParseUint(m.ChannelID, 0, 64)
	if err != nil {
		log.Printf("Error parsing channel ID: %s", err)
		return
	}

	message, err := strconv.ParseUint(m.ID, 0, 64)
	if err != nil {
		log.Printf("Error parsing message ID: %s", err)
		return
	}

	user, err := strconv.ParseUint(m.Author.ID, 0, 64)
	if err != nil {
		log.Printf("Error parsing user ID: %s", err)
		return
	}

	image := model.Image{
		Hash:    hashBase64,
		Guild:   guild,
		Channel: channel,
		Message: message,
		User:    user,
	}

	matches, err := image.FindMatch()
	if err != nil {
		log.Printf("Error finding matches: %s", err)
		return
	}

	if len(matches) != 0 {
		found(s, m, matches)
	}

	err = image.Insert()
	if err != nil {
		log.Printf("Error inserting image: %s", err)
		return
	}
}
