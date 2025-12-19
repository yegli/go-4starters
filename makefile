.PHONY: all tidy test build run install clean

all: tidy test build run install

tidy:
	cd ./greetings && go mod tidy
	cd ./hello && go mod tidy

test:
	cd ./greetings && go test -v

build:
	cd ./hello && go build

run:
	cd ./hello && go run .

install:
	cd ./hello && go install

clean:
	cd ./hello && go clean