# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFLAGS = -ldflags "-s -w" 

all: test

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
