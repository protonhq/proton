.PHONY: build

build:
	@go build -o build/proton cli/main.go

up:
	@docker-compose up