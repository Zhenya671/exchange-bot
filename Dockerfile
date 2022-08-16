FROM golang:1.19-alpine3.15

COPY . /github.com/Zhenya671/telegram-bot-exchangeRates
WORKDIR /github.com/Zhenya671/telegram-bot-exchangeRates
RUN ls -a
RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/Zhenya671/telegram-bot-exchangeRates/bin/bot .
COPY --from=0 /github.com/Zhenya671/telegram-bot-exchangeRates/config config/

RUN ls -a

#COPY --from=0 /github.com/Zhenya671/telegram-bot-exchangeRates/.env .

EXPOSE 80

CMD ["./bot"]
