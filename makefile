build: 
	go build -o app cmd/main.go

build-and-run: build 
	./app

build-docker:
	docker-compose build

run-dev: 
	docker compose up --watch

run:
	./app

# Test
test:
	go test -v ./...

.PHONY: build run test