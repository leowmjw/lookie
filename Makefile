all: clean test build

clean:
	rm -rfv dist

test:
	go test -v .

build:
	go build -o dist/lookied cmd/server/main.go 

dev:
    echo "Development mode ..."

.PHONY: all clean test build dev
