#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /go/src/app
COPY . .
ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags prod -o $GOPATH/bin/app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
CMD ["./app"]
LABEL Name=quran-twitter-bot Version=1.0
