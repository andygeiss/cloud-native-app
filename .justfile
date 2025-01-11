set dotenv-load

# Compile the Go sources.
build: test
    @go build \
    -ldflags "-s -w" \
    -o ./bin/$(basename $PWD) ./cmd/app/main.go

# Install the binary to the $HOME/bin directory.
install: build
    @cp ./bin/$(basename $PWD) $HOME/bin/$(basename $PWD)

# Run the Go sources.
run:
    @./bin/$(basename $PWD)

# Test the Go sources (Units).
test:
    @go test -v -coverprofile=.coverprofile.out ./internal/app/core/services/...
