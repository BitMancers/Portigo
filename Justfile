build:
	mkdir -p build
	CGO_ENABLED=0 go build -o build ./cmd/portigo/

migrate CMD NAME:
	go run ./cmd/portigo/ migrate {{CMD}} {{NAME}}

run: build
	./build/portigo serve --dev

watch:
  watchexec -e go -r "just run"

clean:
	rm -rf build pb_data
