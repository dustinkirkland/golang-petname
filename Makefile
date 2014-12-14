GO_BUILD=go build
GO_INSTALL=go install
GO_TEST=go test
GO_CLEAN=go clean

all: update petname

update:
	./update.sh

petname: petname.go
	$(GO_BUILD)
	GOPATH=$(shell pwd) $(GO_INSTALL) petname

test: petname.go petname_test.go
	GOROOT= GOPATH=$(shell pwd) $(GO_TEST)

clean:
	$(RM) -rf petname petname.go petname.a petname.py petname.pyc   init/ linux_amd64/ pkg/

.PHONY: all clean test
