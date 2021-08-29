LOCAL_BIN:=$(CURDIR)/bin

.PHONY: deps
deps:
	go get -u github.com/onsi/ginkgo
	go get -u github.com/onsi/gomega
	go get -u github.com/golang/mock
	go get -u github.com/rs/zerolog/log
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/rs/zerolog/log
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: generate
make generate:
	protoc --proto_path=. -I vendor.protogen \
	--go_out=pkg/api --go_opt=paths=import \
	--go-grpc_out=pkg/api --go-grpc_opt=paths=import \
	api/api.proto


build:
	go build -o bin/main cmd/ova-track-api/main.go

run:
	go run cmd/ova-track-api/main.go

test:
	go test  ./...

