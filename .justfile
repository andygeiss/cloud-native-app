set dotenv-load

# Compile the Go sources.
build:
    @go build \
    -ldflags "-X main.version=$(git describe --tags --always --dirty)" \
    -o ./bin/$(basename $PWD) ./cmd/app/main.go

# Test the Go sources (Units).
test:
    @go test -v -coverprofile=.coverprofile.out ./internal/app/...
