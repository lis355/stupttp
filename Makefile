APP_NAME = stupttp

.PHONY: clear build build-platform build-linux build-macos

clear:
	rm -rf ./build

build-platform:
	mkdir -p build
	buildfile=./build/$(APP_NAME)-$(PLATFORM)-$(ARCH); \
	GOOS=$(PLATFORM) GOARCH=$(ARCH) go build -ldflags="-s -w" -trimpath -o $$buildfile main.go; \
	du -h $$buildfile

build-linux:
	$(MAKE) build-platform PLATFORM=linux ARCH=amd64

build-macos:
	$(MAKE) build-platform PLATFORM=darwin ARCH=arm64

build: clear build-linux build-macos
