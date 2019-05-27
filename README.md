# Quran Twitter Bot

[![Build Status](https://travis-ci.com/airani/quran-twitter-bot.svg?branch=master)](https://travis-ci.com/airani/quran-twitter-bot)

This is a Quran Twitter bot that tweets every hour an Aye of Quran with Persian translation. This bot wrriten by GoLang.

This bot now tweet to [@HourQuran](https://twitter.com/HourQuran) account at Twitter.

## Build

``` sh
# Download modules
go mod download

# Build
go build -o $GOPATH/bin/quran-twitter-bot

# Run with simple configuration
go run main.go tweetAye --config config/.cobra.json

# Run after build
quran-twitter-bot tweetAye --config config/.cobra.json
```

## TODO

- [x] Add dev and prod environments with config files
- [x] Add git repo
- [x] Add readme file
- [x] Dockerize
- [x] Add test file