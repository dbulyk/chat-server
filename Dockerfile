FROM golang:1.23.2-alpine3.20 AS builder

COPY . /github.com/dbulyk/chat_server/source
WORKDIR /github.com/dbulyk/chat_server/source

RUN go mod download
RUN go build -o ./bin/chat_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/dbulyk/chat_server/source/bin/chat_server /root/

ENTRYPOINT ["./chat_server"]
CMD []