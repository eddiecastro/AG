SHELL := /bin/bash
GO111MODULE := on

.PHONY: all
all: help

.PHONY: build
build: ## build the current go project
	cd ./cmd/salesloft && go build -v main.go && cd ..

.PHONY: run
run: ## run the current go project
	cd ./cmd/salesloft && go run main.go && cd ..