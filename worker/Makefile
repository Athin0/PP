test:
	go test ./...

grammar-gen:
	@echo "==> generating parser"
	rm -rf ./grammar
	gogll -o ./grammar ./Math.md

build:
	rm -rf bin
	go build -tags dynamic -mod=vendor -o ./bin/worker cmd/main.go

build-ci:
	rm -rf bin
	go build -tags=dynamic,musl -mod=mod -o ./bin/worker cmd/main.go

run:
	make build
	./bin/worker