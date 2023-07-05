# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=todo
BINARY_MAC=$(BINARY_NAME)_mac
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WIN=$(BINARY_NAME)_unix

all: test build
build: 
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o ./bin/macos/$(BINARY_NAME) ./cmd/todo
test: 
	$(GOTEST) ./cmd/todo ./...
clean: 
	$(GOCLEAN)
	rm -f ./bin/macos/$(BINARY_NAME)
	rm -f ./bin/linux/$(BINARY_NAME)
	rm -f ./bin/windows/$(BINARY_NAME).exe
run:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/todo ./...
	./$(BINARY_NAME)


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/linux/$(BINARY_NAME) ./cmd/todo
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./bin/windows/$(BINARY_NAME) ./cmd/todo
build-all:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o ./bin/macos/$(BINARY_NAME) ./cmd/todo
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/linux/$(BINARY_NAME) ./cmd/todo
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./bin/windows/$(BINARY_NAME).exe ./cmd/todo