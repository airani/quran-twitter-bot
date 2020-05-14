package bot

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/airani/quran"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/mavihq/persian"
	"github.com/spf13/viper"
)

type Tweet struct {
	Quran   quran.Quran
	QuranFa quran.Quran
}

func (t Tweet) PostRandAyah() error {
	s := t.Quran.RandSurah()
	a := s.RandAyah()

	if a.IsSajdaObligatory() {
		return errors.New("aye is sajda obligatory")
	}

	err := t.Post(t.String(s, a))
	return err
}

func (t Tweet) Post(ts string) error {
	if !t.Valid(ts) {
		return errors.New("aye length more than 280 char (twitter limit) and can't be tweet")
	}
	_, _, err := t.client().Statuses.Update(ts, nil)
	return err
}

func (t Tweet) Valid(s string) bool {
	if utf8.RuneCountInString(s) > maxTweetLen {
		return false
	}
	return true
}

func (t Tweet) String(s quran.Surah, a quran.Ayah) string {
	faS := t.QuranFa.Surah(s.Index).Ayah(a.Index)
	return fmt.Sprintf(
		"«%s»\n\n%s\n\n%s:%s",
		a.Text,
		faS.Text,
		s.Name,
		persian.ToPersianDigitsFromInt(s.Index),
	)
}

func (t Tweet) client() *twitter.Client {
	cOauth1 := oauth1.NewConfig(
		viper.GetString("twitter.CONSUMER_KEY"),
		viper.GetString("twitter.CONSUMER_SECRET_KEY"),
	)

	tokenOauth1 := oauth1.NewToken(
		viper.GetString("twitter.ACCESS_TOKEN"),
		viper.GetString("twitter.ACCESS_TOKEN_SECRET"),
	)

	return twitter.NewClient(cOauth1.Client(oauth1.NoContext, tokenOauth1))
}
