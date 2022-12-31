pkgs  = $(shell GOFLAGS=-mod=mod go list ./... | grep -vE -e /vendor/ -e /pkg/swagger/)


swagger-validate:
	@echo "==> validating swagger declaration"
	swagger validate "${PWD}/pkg/swagger/swagger.yml"


swagger-doc:
	@echo "==> generating swagger doc"
	rm -rf ./doc
	mkdir ./doc
	python3 ./pkg/swagger/swagger-yaml-to-html.py < ./pkg/swagger/swagger.yml > ./doc/index.html


swagger-gen:
	@echo "==> generating swagger go code"
	rm -rf ./internal/generated
	GOFLAGS=-mod=mod go generate ${pkgs}

grammar-gen:
	@echo "==> generating parser"
	rm -rf ./internal/grammar
	gogll -o ./internal/grammar ./pkg/grammar/Math.md

build:
	@echo "==> building application"
	rm -rf bin
	go build -tags=dynamic -mod=vendor -o bin/app cmd/main.go

build-ci:
	@echo "==> building application ci"
	rm -rf bin
	go build -tags=dynamic,musl -mod=vendor -o bin/app cmd/main.go

run:
	@echo "==> running application"
	make build
	./bin/app