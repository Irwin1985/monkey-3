lint:
	golint -set_exit_status $$(go list ./... | grep -v /vendor/)
	go vet ./...

.PHONY: test
test:
	go test ./... -race -cover
