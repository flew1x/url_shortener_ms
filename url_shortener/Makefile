.PHONY: lint build build-docker run run-docker test clean clean-docker

.DEFAULT_GOAL := build

lint:
	golangci-lint run

build:
	go build -o ./.bin/server cmd/app/main.go

build-docker:
	docker build -t url_shortener_auth_ms .

run: build
	./.bin/server

run-docker: build-docker
	docker run -p 8080:8080 -d url_shortener_auth_ms

test:
	go test -v ./...

clean:
	rm -rf ./.bin

clean-docker:
	docker rmi url_shortener_auth_ms