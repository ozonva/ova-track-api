echLOCAL_BIN:=$(CURDIR)/bin
DBSTRING:="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5434/$(POSTGRES_DB)?sslmode=disable"


.PHONY: deps
deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
    GOBIN=$(LOCAL_BIN) go install github.com/golang/protobuf/proto
    GOBIN=$(LOCAL_BIN) go install github.com/golang/protobuf/protoc-gen-go
    GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc
    GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
    GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose
    GOBIN=$(LOCAL_BIN) go install github.com/prometheus/client_golang/prometheus


.PHONY: generate
generate:
	GOBIN=$(LOCAL_BIN) protoc --proto_path=. -I vendor.protogen \
          --go_out=pkg/api --go_opt=paths=import \
          --go-grpc_out=pkg/api --go-grpc_opt=paths=import \
          api/api.proto


vendor-proto:
	mkdir -p vendor.protogen
	mkdir -p vendor.protogen/api/ova-track-api
	cp api/api.proto vendor.protogen/api/api.proto
	@if [ ! -d vendor.protogen/google ]; then \#
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi

proto: vendor-proto generate

.PHONY: migrate-up
migrate-up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir migration status
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir migration up

build:
	go build -o bin/main cmd/ova-track-api/main.go

.PHONY: run
run:
	GOBIN=$(LOCAL_BIN) go run cmd/ova-track-api/main.go

test:
	go test  ./...

