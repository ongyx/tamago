GO=go
MAIN=frontend/main.go
BINARY=tamago
BINARY_WIN=tamago.exe
BINARY_DIR=build

.PHONY: build

all: build
build:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY_DIR)/$(BINARY) $(MAIN)
build-win:
	GOOS=windows GOARCH=amd64 $(GO) build -o $(BINARY_DIR)/$(BINARY_WIN) $(MAIN)
