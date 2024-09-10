.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./input.css -o ./static/css/style.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: dev
dev:
	go build -o ./tmp/ ./cmd/main.go
	air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -o ./bin/ ./cmd/main.go

OS=$(shell uname | tr '[:upper:]' '[:lower:]' | sed 's/darwin/macos/;s/linux/linux/;s/cygwin/windows/;s/mingw/windows/')
ARCH=$(shell uname -m | sed 's/x86_64/x64/;s/aarch64/arm64/')
TWCLI_FILENAME=tailwindcss-$(OS)-$(ARCH)

.PHONY: tailwind-update
tailwind-update:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/$(TWCLI_FILENAME)
	chmod +x $(TWCLI_FILENAME)
	mv $(TWCLI_FILENAME) tailwindcss

