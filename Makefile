.PHONY: all build run clean

build:
	mkdir -p build
	CGO_ENABLED=0 go build -o build ./cmd/portigo/

run: build
	./build/portigo serve --dev

clean:
	rm -rf build pb_data
