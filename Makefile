# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary
BINARY=go-web-skeleton

# Commands
all: test build
build:
	$(GOBUILD) -o $(BINARY) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY)
run:
	$(GOBUILD) -o $(BINARY) -v ./...
	./$(BINARY)
