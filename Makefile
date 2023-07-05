# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=todo
BINARY_MAC=$(BINARY_NAME)_mac
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WIN=$(BINARY_NAME)_win.exe

all: test build
build: 
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o ./bin/$(BINARY_MAC) ./cmd/todo
test: 
	$(GOTEST) ./cmd/todo ./...
clean: 
	$(GOCLEAN)
	rm -f ./bin/$(BINARY_MAC)
	rm -f ./bin/$(BINARY_UNIX)
	rm -f ./bin/$(BINARY_WIN)
run:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/todo ./...
	./$(BINARY_NAME)


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/$(BINARY_UNIX) ./cmd/todo
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./bin/$(BINARY_WIN) ./cmd/todo
build-all:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o ./bin/$(BINARY_MAC) ./cmd/todo
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/$(BINARY_UNIX) ./cmd/todo
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./bin/$(BINARY_WIN) ./cmd/todo