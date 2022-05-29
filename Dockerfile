FROM golang:1.17.8 as builder

MAINTAINER cunoe

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/go-wows-bot

COPY . .

RUN  go build -o app

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/go-wows-bot/app .
COPY --from=builder /go/src/go-wows-bot/template/ ./template/

CMD ["./app"]

