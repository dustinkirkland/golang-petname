GO_BUILD=go build
GO_TEST=go test
GO_CLEAN=go clean

all: update petname

update:
	./update.sh

petname: petname.go
	$(GO_BUILD) -o $@ $<

test: petname.go petname_test.go
	$(GO_TEST)

clean:
	$(RM) petname petname.go

.PHONY: all clean test
