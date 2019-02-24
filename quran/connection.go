package quran

import (
	"github.com/spf13/viper"
	"upper.io/db.v3/mongo"
)

func Mongo() mongo.ConnectionURL {
	return mongo.ConnectionURL{
		Database: viper.GetString("mongodb.database"),
		Host:     viper.GetString("mongodb.host"),
		User:     viper.GetString("mongodb.user"),
		Password: viper.GetString("mongodb.password"),
	}
}
