.PHONY: all clean

binary   = gupi
version  = 0.0.1
build	   = $(shell git rev-parse HEAD)
ldflags  = -ldflags "-X 'github.com/phantompunk/gupi/command.version=$(version)'
ldflags += -X 'github.com/phantompunk/gupi/command.build=$(build)'"

all:
	go build -o $(binary) $(ldflags)

test:
	go test ./... -cover -coverprofile c.out
	go tool cover -html=c.out -o coverage.html

clean:
	rm -rf $(binary) c.out coverage.html