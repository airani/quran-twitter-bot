# Quran Twitter Bot

This is a Quran Twitter bot that tweets every hour an Aye of Quran with Persian translation. This bot wrriten by GoLang.

This bot now tweet to [@HourQuran](https://twitter.com/HourQuran) account at Twitter.

## Build

To build this go app first you need create config file for your environment from `config/config_sample.go` like `config/config_dev.go` or `config/config_prod.go` then you should execute build command with tags like below:

``` sh
go build -tags prod -o $GOPATH/bin/quran-twitter-bot
```

## TODO

- [x] Add dev and prod environments with config files
- [x] Add git repo
- [x] Add readme file
- [ ] Dockerize
- [ ] Add test file