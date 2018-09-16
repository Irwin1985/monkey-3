lint:
	golint -set_exit_status $$(go list ./... | grep -v /vendor/)
	go vet ./...

.PHONY: test
test:
	go test ./... -race -cover

build:
	go build ./cmd/monkey

install:
	go install ./cmd/monkey

clean:
	rm -f monkey
	go clean
