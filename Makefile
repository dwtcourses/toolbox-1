build:
	./scripts/build.sh

test:
	GO111MODULE=on go test ./...

run: build
	./toolbox
