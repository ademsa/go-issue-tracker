default: build

clean:
	rm -r ./bin && mkdir ./bin

prepare-dependencies:
	go mod download

build-code:
	go build -o ./bin/go-issue-tracker ./cmd/main.go

build: clean prepare-dependencies build-code

run-code:
	./bin/go-issue-tracker

build-and-run: build run-code

test:
	go test -v ./... -coverprofile=coverage.out
 
test-coverage:
	go tool cover -func=coverage.out

test-coverage-html:
	go tool cover -html=coverage.out

test-coverage-html-file:
	go tool cover -html=coverage.out -o coverage.html

test-and-analyze: test test-coverage test-coverage-html-file