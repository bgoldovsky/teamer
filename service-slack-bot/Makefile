.PHONY:build
build:
	echo "building.."
	go build -o bin/service-slack-bot ./cmd/service/main.go

.PHONY:run
run:
	echo "running.."
	go run -race ./cmd/service/main.go

.PHONY:test
test:
	echo "testing.."
	go clean -testcache
	go test -v -cover ./...

.PHONY:proto
proto:
	protoc -I proto --go_out=plugins=grpc,paths=import:. proto/*.proto

.DEFAULT_GOAL := run