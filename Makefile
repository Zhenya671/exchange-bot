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

remove:
	docker rm `docker ps -a -q -l`

start-container-test:
	docker run --name exchange-rates-bot --env TOKEN=5558699531:AAG7HFzLoxQf7QQhC-VsGj_yFZCf7o9XW9I telegram-bot