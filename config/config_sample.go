// +build sample

package config

import "upper.io/db.v3/mongo"

const (
	ENV                         = "sample"
	TWITTER_CONSUMER_KEY        = "xxx"
	TWITTER_CONSUMER_SECRET_KEY = "yyy"
	TWITTER_ACCESS_TOKEN        = "123-zzz"
	TWITTER_ACCESS_TOKEN_SECRET = "xyz"
)

func Mongo() mongo.ConnectionURL {
	return mongo.ConnectionURL{
		Database: `quran`,
		Host:     `localhost`,
		User:     `usr`,
		Password: `123456`,
	}
}
