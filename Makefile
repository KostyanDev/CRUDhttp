include .env
export

default: help

## build: Building server.
build:
	@echo "***** Building... *****"
	@go build -o server cmd/main.go
	@echo "Success!"

## start: Start server.
start: build
	@echo "***** Start server... *****"
	@echo "LINK -> http://localhost:ma$(PORT) <-"
	@./server

## clean: Remove server.exe
clean:
	@echo "***** Clean... *****"
	rm -rf server
	@echo "Cleared!"

help: Makefile
	@echo "Choose a command run in "HTTPCrud". Example: make [command]"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'