.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t telegram-bot .

start-container:
	docker run --name exchange-rates-bot --env-file .env telegram-bot

#detached mod
continue:
	docker start `docker ps -q -l`

stop:
	docker stop `docker ps -q -l`