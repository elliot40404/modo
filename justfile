set windows-shell := ["pwsh.exe", "-NoLogo", "-Command"]

default: build

build_cmd := if os() == "windows" { "go build -o ./bin/modo.exe ./cmd/modo/" } else { "go build -o ./bin/modo ./cmd/modo/" }

build: clean lint
    {{build_cmd}}

run +args:
    go run ./cmd/modo/ {{args}}

install:
    go install ./cmd/modo/

build-run +args: build
    ./bin/modo {{args}}

rmcmd := if os() == "windows" { "mkdir ./bin -Force; Remove-Item -Recurse -Force ./bin" } else { "rm -rf ./bin" }

clean:
    {{rmcmd}}

lint:
    golangci-lint run

lint-fix:
    golangci-lint run --fix

vendor:
    go mod tidy
    go mod vendor
    go mod tidy

release:
    goreleaser release --snapshot --clean
