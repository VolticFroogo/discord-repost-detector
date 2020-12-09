package model

import (
	"fmt"
	"os"
)

var (
	DiscordToken = os.Getenv("DISCORD_TOKEN")
	DiscordID    = os.Getenv("DISCORD_ID")
	Ping         = fmt.Sprintf("<@!%s>", DiscordID)
)
