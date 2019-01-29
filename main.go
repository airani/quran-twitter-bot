package main

import (
	"errors"
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

func main() {
	ticker := time.NewTicker(1 * time.Hour)

	var aye Aye
	var err error

	for range ticker.C {
		for {
			aye, err = newAyeByRand()
			if err != nil {
				LogToFile(fmt.Sprintf("%q", err))
				break
			}

			err := aye.sendAsTweet()
			if err != nil {
				LogToFile(fmt.Sprintf("%q", err))
			} else {
				break
			}

		}
	}
}

const (
	mongoDbAyeColl  = "aye"
	mongoDbSuraColl = "sura"

	logFile = "quran-tweet-bot_error.log"
)

// Aye strcut
type Aye struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	SuraID     bson.ObjectId `bson:"_sura_id,omitempty"`
	Number     uint          `bson:"number,omitempty"`
	Text       string        `bson:"text,omitempty"`
	SimpleText string        `bson:"simple_text,omitempty"`
	Translate  `bson:"translate"`
	Sura
}

// RandAye returns randomly an Aye from Quran
func newAyeByRand() (aye Aye, err error) {
	sess, err := mongo.Open(config.Mongo())
	if err != nil {
		return
	}
	defer sess.Close()

	res := sess.Collection(mongoDbAyeColl).Find().
		Limit(1).
		Offset(rand.Intn(6236))

	err = res.One(&aye)
	if err != nil {
		return
	}

	var sura Sura
	err = sess.Collection(mongoDbSuraColl).
		Find(db.Cond{"_id": aye.SuraID}).
		One(&sura)

	aye.Sura = sura

	return
}

// FormatAye to prepare as string for tweet
func (a *Aye) FormatAye() string {
	return fmt.Sprintf("«%s»\n\n%s\n\n%s:%s",
		a.Text,
		a.Translate.FooladvandFa,
		a.Sura.Name,
		persian.ToPersianDigitsFromInt(int(a.Number)))
}

// CanTweet check a string can be tweet or not by checking string length
func (a *Aye) CanTweet() bool {
	s := a.FormatAye()
	if utf8.RuneCountInString(s) > 280 {
		return false
	}

	return true
}

func (a *Aye) sendAsTweet() error {
	if !a.CanTweet() {
		return errors.New("can't send tweet")
	}
	configOauth1 := oauth1.NewConfig(config.TWITTER_CONSUMER_KEY, config.TWITTER_CONSUMER_SECRET_KEY)
	tokenOauth1 := oauth1.NewToken(config.TWITTER_ACCESS_TOKEN, config.TWITTER_ACCESS_TOKEN_SECRET)
	httpClient := configOauth1.Client(oauth1.NoContext, tokenOauth1)
	client := twitter.NewClient(httpClient)
	_, _, err := client.Statuses.Update(a.FormatAye(), nil)

	return err
}

// Translate struct
type Translate struct {
	FooladvandFa string `bson:"fa-fooladvand"`
	MakaremFa    string `bson:"fa-makarem"`
	GhomesheiFa  string `bson:"fa-ghomshei"`
}

// Sura struct
type Sura struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Number uint          `bson:"number,omitempty"`
	Name   string        `bson:"name,omitempty"`
	Ayat   uint          `bson:"cnt_aye"`
}

// LogToFile write log to a file
func LogToFile(s string) {
	f, err := os.OpenFile(logFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(s)
	log.Println(s)
}
