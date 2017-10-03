NAME := abayo

all: build

setup:
	go get -u github.com/aws/aws-sdk-go

build:
	go build -o bin/$(NAME)

clean:
	rm -rf bin/*
