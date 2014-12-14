GO_BUILD=go build
GO_INSTALL=go install
GO_TEST=go test
GO_CLEAN=go clean

all: update petname

update:
	./update.sh

petname: petname.go cmd/petname/main.go
	$(GO_BUILD)
	GOPATH=$(shell pwd) $(GO_BUILD) -o cmd/petname/petname cmd/petname/main.go

test: petname.go petname_test.go
	GOPATH=$(shell pwd) $(GO_TEST)

clean:
	$(RM) -rf petname petname.pyc

.PHONY: all clean test
