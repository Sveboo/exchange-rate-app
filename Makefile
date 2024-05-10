GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=currency_service

build:
	$(GOCMD) mod tidy
	$(GOBUILD) -o $(BINARY_NAME) cmd/main.go

clean:
	rm -f $(BINARY_NAME)

default: build