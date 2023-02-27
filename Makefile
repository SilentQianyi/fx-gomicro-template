GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/envoyproxy/protoc-gen-validate@latest
	@go install github.com/golang/mock/mockgen@main
#	@go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
#	@go install github.com/jackc/tern@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=.. --go_out=:.. proto/common/v1/common.proto
	@protoc --proto_path=. --micro_out=.. --go_out=:.. proto/hello_world/v1/hello_world.proto
	@#protoc --proto_path=. --micro_out=. --go_out=:. proto/health.proto
	

.PHONY: api
api:
	@/bin/sh ./scripts/api_gen.sh false

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o chat *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@DOCKER_BUILDKIT=1 docker build -t chat:latest .

.PHONY: sqlc
sqlc:
	@sqlc generate -f ./postgres/sqlc.yaml