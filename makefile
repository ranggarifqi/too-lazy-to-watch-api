build: 
	go build -o app cmd/main.go

run: build 
	./app

build-docker:
	docker-compose build

run-docker: 
	docker-compose up

# Test
test:
	go test -v ./...

.PHONY: build run test