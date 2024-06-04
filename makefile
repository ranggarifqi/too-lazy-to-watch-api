build: 
	go build -o app cmd/main.go

run: build 
	./app

# Test
test:
	go test -v ./...

.PHONY: build run test