.PHONY:build
build:
	echo "building.."
	go build -o bin/service-duty-bot ./cmd/service/main.go

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
	protoc --go_out=. data-bus/proto/*.proto

.DEFAULT_GOAL := run