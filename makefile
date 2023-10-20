GOBIN ?= $$(go env GOPATH)/bin

.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: check-coverage
check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.github/.testcoverage.yml

# Runs tests on entire repo
.PHONY: test
test: 
	go test -timeout=3s -race -count=10 -failfast -short ./...
	go test -timeout=3s -race -count=1 -failfast ./...