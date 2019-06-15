DOCKER_TARGET?=ci

setup:
	go run github.com/google/wire/cmd/wire
	go mod vendor
	echo "To start server: make start"

start: wire
	go run cmd/main.go

wire:
	cd pkg && go generate

test:
	cd pkg && go test
	
cleanup:
	rm -rf ./build
	rm -rf ./vendor

default:
	setup