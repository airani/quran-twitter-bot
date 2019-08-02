package bot

import (
	"errors"
	"fmt"
	"log"
	"time"
	"unicode/utf8"

	"github.com/airani/quran-twitter-bot/quran"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/mavihq/persian"
	"github.com/spf13/viper"
)

const (
	maxTweetLen int           = 280
	interval    time.Duration = 1 * time.Hour
)

// Run twitter bot
func Run() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Quran twitter bot started..")

	ticker := time.NewTicker(interval)

	for range ticker.C {
		for {
			err := tweetRandAya()
			if err != nil {
				log.Println(fmt.Sprintf("%q", err))
			} else {
				break
			}

		}
	}
}

func tweetRandAya() error {
	q := quran.New()
	s := q.RandSura()
	a := s.RandAya()

	if a.IsSajdaObligatory() {
		return errors.New("aye is sajda obligatory")
	}

	ts := tweetAyeText(s, a)
	err := tweet(ts)
	return err
}

func tweet(ts string) error {
	if !canTweet(ts) {
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
	_, _, err := client.Statuses.Update(ts, nil)

	return err
}

func canTweet(s string) bool {
	if utf8.RuneCountInString(s) > maxTweetLen {
		return false
	}
	return true
}

func tweetAyeText(s quran.Sura, a quran.Aya) string {
	faS := quran.Fa().Sura(s.Index).Aya(a.Index)
	return fmt.Sprintf(
		"«%s»\n\n%s\n\n%s:%s",
		a.Text,
		faS.Text,
		s.Name,
		persian.ToPersianDigitsFromInt(s.Index),
	)
}
