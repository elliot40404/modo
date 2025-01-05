default: build

build: clean lint
    go build -o ./bin/modo ./cmd/modo/

run +args:
    go run ./cmd/modo/ {{args}}

install:
    go install ./cmd/modo/

build-run +args: build
    ./bin/modo {{args}}

clean:
    rm -rf ./bin

lint:
    golangci-lint run

vendor:
    go mod tidy
    go mod vendor
    go mod tidy

release:
    goreleaser release --snapshot --clean
