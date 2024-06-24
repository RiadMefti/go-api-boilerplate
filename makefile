.PHONY: build run

all: run


build:
	@go build -o bin/out cmd/main.go

run: build
	@./bin/out
