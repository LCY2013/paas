
.PHONY: proto
proto:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) luochunyun/protoc:1.0.0 --proto_path=. --micro_out=. --go_out=:. ./proto/svcapi/svcapi.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o svcapi *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t svcapi:latest

.PHONY: build-in-docker
build-in-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o svcapi *.go
	docker build . -t luochunyun/svcapi:latest
	rm -rf svcapi
