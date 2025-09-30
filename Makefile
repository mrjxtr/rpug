.PHONY: build build_amd64 run clean_all clean test tidy all

# Application name
APP_NAME := rpug

# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOMOD := $(GOCMD) mod
GOGET := $(GOCMD) get

# Main package path
MAIN_PATH := ./main.go
BIN_PATH := ./bin

$(BIN_PATH):
	mkdir -p $@

#Build application for amd64
build_amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v -o $(BIN_PATH)/$(APP_NAME)-linux-x86_64 $(MAIN_PATH)

#Build application
build:
	$(GOBUILD) -v -o $(APP_NAME) $(MAIN_PATH)

# Run the application
run:
	$(GORUN) $(MAIN_PATH)

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BIN_PATH)/$(APP_NAME)*
	rm -f $(APP_NAME)*

# Clean all build artifacts
clean_all:
	$(GOCLEAN)
	rm -f $(BIN_PATH)/$(APP_NAME)*
	rm -f $(APP_NAME)*

# Run tests
test:
	$(GOTEST) -v ./...

# Update dependencies
tidy:
	$(GOMOD) tidy

# Default target
all: clean build 
