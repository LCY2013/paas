
.PHONY: proto
proto:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) luochunyun/protoc:1.0.0 --proto_path=. --micro_out=. --go_out=:. ./proto/middleware/middleware.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o middleware *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t middleware:latest

.PHONY: build-in-docker
build-in-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o middleware *.go
	docker build . -t luochunyun/middleware:latest
	rm -rf middleware
