.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o tool *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t tool:latest

.PHONY: build-in-docker
build-in-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o tool *.go
	docker build . -t luochunyun/tool:latest
	rm -rf tool