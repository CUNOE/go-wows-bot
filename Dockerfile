FROM golang:1.17.8 as builder

MAINTAINER cunoe

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/go-wows-bot

COPY . .

RUN  go build

FROM ubuntu:20.04 as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/go-wows-bot/go-wows-bot .
COPY --from=builder /go/src/go-wows-bot/template/ ./template/

CMD ["./go-wows-bot"]
