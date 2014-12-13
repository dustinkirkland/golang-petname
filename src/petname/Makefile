GO_BUILD=go build
GO_INSTALL=go install
GO_TEST=go test
GO_CLEAN=go clean

all: update petname

update:
	./update.sh

petname: petname.go
	$(GO_BUILD) petname.go
	GOPATH=$(shell pwd)/../../ $(GO_INSTALL) petname

test: petname.go petname_test.go
	$(GO_TEST)

clean:
	$(RM) petname petname.go petname.a

.PHONY: all clean test
