.PHONY: lint build build-docker run run-docker test clean clean-docker

.DEFAULT_GOAL := build

lint:
	golangci-lint run

build:
	go build -o ./.bin/server cmd/app/main.go

build-docker:
	docker build -t url_shortener_ms .

run: build
	./.bin/server

run-docker:
	cd .. && docker compose down && docker-compose up --build --force-recreate

test:
	go test -v ./...

test-cover:
	go test -coverprofile=coverage.out -covermode=atomic ./...

clean:
	rm -rf ./.bin

clean-docker:
	docker rmi url_shortener_auth_ms

MOCKS_DESTINATION=mocks

.PHONY: mocks

mocks: $(wildcard internal/service/*.go) $(wildcard internal/repository/*.go) $(wildcard internal/cache/*.go) $(wildcard internal/config/*.go) $(wildcard internal/entity/*.go)
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@mkdir -p $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$(basename $$file) -package=mocks; done

clean-mocks:
	rm -f $(MOCKS_DIR)/*.go