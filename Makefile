build:
	go build -o tarantula -v

clean:
	go clean
	rm -f tarantula

run: build
	./tarantula

install:
	go get -u github.com/keybase/go-keybase-chat-bot/...
