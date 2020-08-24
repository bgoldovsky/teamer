.PHONY:compose
compose:
	 docker-compose build && docker-compose up

.PHONY:build
build:
	 docker-compose build

.PHONY:up
up:
	 docker-compose up

.PHONY:test
test:
	echo "testing.."
	go clean -testcache
	go test -v -cover ./...

.DEFAULT_GOAL := compose