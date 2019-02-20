package quran

import (
	"fmt"
	"log"
	"time"
	"unicode/utf8"
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

// CanTweet check a string can be tweet or not by checking string length
func CanTweet(s string) bool {
	if utf8.RuneCountInString(s) > maxTweetLen {
		return false
	}

	return true
}
