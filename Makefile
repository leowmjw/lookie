all: clean test build
clean:
	rm -rfv dist
test:
	go test -v .
build:
	go build -o dist/lookied cmd/server/main.go 
.PHONY: all clean test build
