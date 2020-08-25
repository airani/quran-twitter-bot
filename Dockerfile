#build stage
FROM golang:1.14-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /go/src/github.com/airani/quran-twitter-bot
COPY go.mod $WORKDIR
COPY go.sum $WORKDIR
ENV GO111MODULE=on
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags prod -o $GOPATH/bin/app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add tzdata
RUN ls /usr/share/zoneinfo
RUN cp /usr/share/zoneinfo/Asia/Tehran /etc/localtime
RUN echo "Asia/Tehran" >  /etc/timezone
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/github.com/airani/quran-twitter-bot/dataset /dataset
COPY --from=builder /go/src/github.com/airani/quran-twitter-bot/config /config
ENV PATH_DATASET="/dataset/"
CMD ["./app", "tweetAye", "--config=/config/.prod.json"]
LABEL Name=quran-twitter-bot Version=1.0
