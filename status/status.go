package status

import (
	"fmt"
	"github.com/VolticFroogo/discord-repost-detector/db"
	"github.com/VolticFroogo/discord-repost-detector/discord"
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

func Start() (quit chan bool, finished chan bool) {
	quit = make(chan bool)
	finished = make(chan bool)

	go thread(quit, finished)
	log.Println("Started status updater thread.")

	return
}

func thread(quit chan bool, finished chan bool) {
	ticker := time.NewTicker(time.Minute)

	updateStatus()

	for {
		select {
		case <-quit:
			log.Println("Ending status updater thread.")
			finished <- true
			return

		case <-ticker.C:
			updateStatus()
		}
	}
}

func updateStatus() {
	defer sentry.Recover()

	count, err := db.Images.EstimatedDocumentCount(db.DefaultContext())
	if err != nil {
		panic(err)
	}

	err = discord.Discord.UpdateStatus(0, fmt.Sprintf("@me - seen %d images", count))
	if err != nil {
		panic(err)
	}
}
