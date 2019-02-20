package quran

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/globalsign/mgo/bson"
	"github.com/mavihq/persian"
	"github.com/spf13/viper"
	db "upper.io/db.v3"
	"upper.io/db.v3/mongo"
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

// newAyeByRand returns randomly an Aye from Quran
func newAyeByRand() (aye Aye, err error) {
	sess, err := mongo.Open(Mongo())
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

// FormattedAsTweet to prepare as string for tweet
func (a *Aye) FormattedAsTweet() string {
	return fmt.Sprintf("«%s»\n\n%s\n\n%s:%s",
		a.Text,
		a.Translate.FooladvandFa,
		a.Sura.Name,
		persian.ToPersianDigitsFromInt(int(a.Number)))
}

func (a *Aye) sendAsTweet() error {
	if !CanTweet(a.FormattedAsTweet()) {
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
	_, _, err := client.Statuses.Update(a.FormattedAsTweet(), nil)

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