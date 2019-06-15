DOCKER_TARGET?=ci

setup:
	go get github.com/google/wire/cmd/wire
	go mod vendor

start:
	go run cmd/main.go

generate: gqlgen wire

wire:
	cd pkg && go generate

test: gen-templates
	cd pkg && go test
	
cleanup:
	rm -rf ./build
	rm -rf ./vendor
	rm -rf ${GOPATH}/bin/nimbus

default:
	setup