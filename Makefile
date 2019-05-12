INSTALLPRE = /usr/local
GOLIST = $(shell go list ./...)
COMMAND_NAME = s3artifact
SUPPORTED_SYSTEMS = linux windows darwin

all: build

build: clean
		@echo "Building binary..."
		go build -o bin/${COMMAND_NAME}

test:
		@test -z "$(gofmt -s -l . | tee /dev/stderr)"
		@test -z "$(golint $(GOLIST) | tee /dev/stderr)"
		@go test -race -test.v $(GOLIST)
		@go vet $(GOLIST)

clean:
		@echo "Cleaning up..."
		rm -rf bin

install:
		@echo "Installing to $(INSTALLPRE)..."
		cp bin/$(COMMAND_NAME) $(INSTALLPRE)/bin/

cross_compile: linux windows darwin

linux:
		@mkdir -p bin/linux
		GOOS=linux go build -v -o bin/linux/$(COMMAND_NAME)

windows:
		@mkdir -p bin/windows
		GOOS=windows go build -v -o bin/windows/$(COMMAND_NAME).exe

darwin:
		@mkdir -p bin/darwin
		GOOS=darwin go build -v -o bin/darwin/$(COMMAND_NAME)
