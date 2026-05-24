.PHONY: migrate all build run clean

build:
	mkdir -p build
	CGO_ENABLED=0 go build -o build ./cmd/portigo/

migrate:
	go run ./cmd/portigo/ migrate $(cmd) $(name)

dev:
	go run ./cmd/portigo serve --dev

run: build
	./build/portigo serve --dev

watch:
	watchexec -e go -r "make dev"

clean:
	rm -rf build pb_data
