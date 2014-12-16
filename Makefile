GO_BUILD=go build
GO_INSTALL=go install
GO_TEST=go test
GO_CLEAN=go clean

all: petname

petname: petname.go cmd/petname/main.go
	$(GO_BUILD)
	GOPATH=$(shell pwd) $(GO_BUILD) -o golang-petname cmd/petname/main.go

test: petname.go petname_test.go
	GOPATH=$(shell pwd) $(GO_TEST)

clean:
	$(RM) -f golang-petname

.PHONY: all clean test
