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
