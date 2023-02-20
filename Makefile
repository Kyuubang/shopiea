include .env

.PHONY: run

build:
	go build -o bin/shopiea

run:
	go run main.go