.PHONY:build
build:
	echo "building.."
	go build -o bin/service-rus-id ./cmd/service/main.go

.PHONY:run
run:
	echo "running.."
	go run -race ./cmd/service/main.go

.PHONY:test
test:
	echo "testing.."
	go clean -testcache
	go test -v -cover -tags=entities ./...

.PHONY:proto
proto:
	protoc -I api/proto --go_out=plugins=grpc,paths=import:. api/proto/*.proto

.PHONY:compose
compose:
	 docker-compose build && docker-compose up

.PHONY:client
client:
	go run ./cmd/client -h 10

.DEFAULT_GOAL := run