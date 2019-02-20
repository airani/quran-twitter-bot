package quran

import (
	"fmt"
	"log"
	"time"
)

const (
	mongoDbAyeColl  string        = "aye"
	mongoDbSuraColl string        = "sura"
	maxTweetLen     int           = 280
	interval        time.Duration = 1 * time.Hour
)

func RunTweetSender() {
	log.Println("Quran twitter bot started..")
	ticker := time.NewTicker(interval)

	for range ticker.C {
		for {
			aye, err := newAyeByRand()
			if err != nil {
				log.Println(fmt.Sprintf("%q", err))
				break
			}

			err = aye.sendAsTweet()
			if err != nil {
				log.Println(fmt.Sprintf("%q", err))
			} else {
				break
			}

		}
	}
}
