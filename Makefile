ROOT=$(shell pwd)

.PHONY: init
init:
	rm -rf build & mkdir -p build/config

.PHONY: build
build: init
	cd ./gcapi && \
	go mod tidy && \
	GOOS=linux GOARCH=amd64 go build -v -o $(ROOT)/build/gcapi main.go && \
	cp config/config.yaml $(ROOT)/build/config && \
	cd $(ROOT) && \
	tar -zcvf backend.tar.gz build
