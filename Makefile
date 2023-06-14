build:
	go build -o  bin/gobank ./cmd/main.go 

run:
	./bin/gobank

test:
	go test -v ./...

