package status

import (
	"fmt"
	"github.com/VolticFroogo/discord-repost-detector/db"
	"github.com/VolticFroogo/discord-repost-detector/discord"
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
	count, err := db.Images.EstimatedDocumentCount(db.DefaultContext())
	if err != nil {
		log.Printf("Error counting image documents: %s", err)
		return
	}

	err = discord.Discord.UpdateStatus(0, fmt.Sprintf("@me - seen %d images", count))
	if err != nil {
		log.Printf("Error updating status: %s", err)
	}
}
