.PHONY: build

build:
	mkdir -p build
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o build/stupttp main.go
	du -h build/stupttp
