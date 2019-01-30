#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
LABEL Name=quran-twitter-bot Version=0.0.1
