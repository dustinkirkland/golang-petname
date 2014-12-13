GO_BUILD=go build
GO_INSTALL=go install
GO_TEST=go test
GO_CLEAN=go clean

all: petname

petname: petname_cmd.go
	GOPATH=$(PWD)/../../ $(GO_BUILD) -o petname petname_cmd.go

clean:
	$(RM) petname

.PHONY: all clean
