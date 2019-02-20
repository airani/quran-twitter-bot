# Quran Twitter Bot

This is a Quran Twitter bot that tweets every hour an Aye of Quran with Persian translation. This bot wrriten by GoLang.

This bot now tweet to [@HourQuran](https://twitter.com/HourQuran) account at Twitter.

## Build

To build this go app first you need install go modules with `go mod`, [read more](https://dev.to/defman/introducing-go-mod-1cdo) about this. Then create config file for your environment from `config/config_sample.go` like `config/config_dev.go` or `config/config_prod.go` then you should execute build command with [build tags](https://golang.org/pkg/go/build/#hdr-Build_Constraints) like below:

``` sh
go mod download

# Run with simple configuration
go run main.go tweetAye --config cmd/.cobra.json

# Build
go build -tags prod -o $GOPATH/bin/quran-twitter-bot


```

## TODO

- [x] Add dev and prod environments with config files
- [x] Add git repo
- [x] Add readme file
- [x] Dockerize
- [ ] Add test file