package quran

import (
	"errors"
	"fmt"
	"log"
	"time"
	"unicode/utf8"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
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

			err = tweet(aye)
			if err != nil {
				log.Println(fmt.Sprintf("%q", err))
			} else {
				break
			}

		}
	}
}

func tweet(a Aye) error {
	if !CanTweet(a.String()) {
		return errors.New("aye length more than 280 char (twitter limit) and can't be tweet")
	}

	configOauth1 := oauth1.NewConfig(
		viper.GetString("twitter.CONSUMER_KEY"),
		viper.GetString("twitter.CONSUMER_SECRET_KEY"),
	)

	tokenOauth1 := oauth1.NewToken(
		viper.GetString("twitter.ACCESS_TOKEN"),
		viper.GetString("twitter.ACCESS_TOKEN_SECRET"),
	)

	httpClient := configOauth1.Client(oauth1.NoContext, tokenOauth1)
	client := twitter.NewClient(httpClient)
	_, _, err := client.Statuses.Update(a.String(), nil)

	return err
}

// CanTweet check a string can be tweet or not by checking string length
func CanTweet(s string) bool {
	if utf8.RuneCountInString(s) > maxTweetLen {
		return false
	}

	return true
}
