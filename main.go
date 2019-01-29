package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"

	"github.com/airani/quran-twitter-bot/config"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/mavihq/persian"
	"gopkg.in/mgo.v2/bson"
	db "upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

type Aye struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	SuraID     bson.ObjectId `bson:"_sura_id,omitempty"`
	Number     uint          `bson:"number,omitempty"`
	Text       string        `bson:"text,omitempty"`
	SimpleText string        `bson:"simple_text,omitempty"`
	Translate  `bson:"translate"`
	Sura
}

type Translate struct {
	FooladvandFa string `bson:"fa-fooladvand"`
	MakaremFa    string `bson:"fa-makarem"`
	GhomesheiFa  string `bson:"fa-ghomshei"`
}

type Sura struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Number uint          `bson:"number,omitempty"`
	Name   string        `bson:"name,omitempty"`
	Ayat   uint          `bson:"cnt_aye"`
}

func main() {
	ticker := time.NewTicker(1 * time.Hour)

	var aye Aye
	var err error

	for range ticker.C {
		for {
			aye, err = RandAye()
			if err != nil {
				LogToFile(fmt.Sprintf("%q", err))
				break
			} else {
				if CanTweet(FormatAye(aye)) {
					err = Tweet(FormatAye(aye))
					if err != nil {
						LogToFile(fmt.Sprintf("%q", err))
					} else {
						break
					}
				}
			}
		}
	}
}

// RandAye returns randomly an Aye from Quran
func RandAye() (Aye, error) {
	var aye Aye

	sess, err := mongo.Open(config.Mongo())
	if err != nil {
		return aye, err
	}
	defer sess.Close()

	ayeColl := sess.Collection("aye")

	res := ayeColl.Find().
		Limit(1).
		Offset(rand.Intn(6236))

	err = res.One(&aye)
	if err != nil {
		return aye, err
	}

	suraColl := sess.Collection("sura")

	var sura Sura
	err = suraColl.Find(db.Cond{"_id": aye.SuraID}).
		One(&sura)

	aye.Sura = sura

	return aye, err
}

// FormatAye to prepare as string for tweet
func FormatAye(aye Aye) string {
	return fmt.Sprintf("«%s»\n\n%s\n\n%s:%s",
		aye.Text,
		aye.Translate.FooladvandFa,
		aye.Sura.Name,
		persian.ToPersianDigitsFromInt(int(aye.Number)))
}

// CanTweet check a string can be tweet or not by checking string length
func CanTweet(s string) bool {
	if utf8.RuneCountInString(s) > 280 {
		return false
	}

	return true
}

// Tweet a string to twitter account
func Tweet(t string) error {
	configOauth1 := oauth1.NewConfig(config.TWITTER_CONSUMER_KET, config.TWITTER_CONSUMER_SECRET_KEY)
	tokenOauth1 := oauth1.NewToken(config.TWITTER_ACCESS_TOKEN, config.TWITTER_ACCESS_TOKEN_SECRET)
	httpClient := configOauth1.Client(oauth1.NoContext, tokenOauth1)
	client := twitter.NewClient(httpClient)
	_, _, err := client.Statuses.Update(t, nil)

	return err
}

// LogToFile write log to a file
func LogToFile(s string) {
	f, err := os.OpenFile("quran-tweet-bot_error.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(s)
	log.Println(s)
}
