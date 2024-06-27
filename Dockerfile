FROM golang:1.22.2 AS chat-api.cache-service-builder

RUN go version

COPY . /github.com/Woodfyn/cache-service/
WORKDIR /github.com/Woodfyn/cache-service/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=chat-api.cache-service-builder /github.com/Woodfyn/cache-service/.bin/app .
COPY --from=chat-api.cache-service-builder /github.com/Woodfyn/cache-service/main.env .
COPY --from=chat-api.cache-service-builder /github.com/Woodfyn/cache-service/configs ./configs

CMD ["./app"]