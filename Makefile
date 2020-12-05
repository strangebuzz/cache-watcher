## —— 🐝 The Cache Watcher Makefile 🐝 —————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Project 🐝 ———————————————————————————————————————————————————————————————
run: ## Run the main go file on a Symfony 5 project
	$(eval path ?= ../strangebuzz.com)
	go run cw.go $(path)

build: ## Build the cw executable for the current platform
	go build cw.go
	shasum -a 256 cw

build-darwin: ## Build for Darwin OS
	GOOS=darwin GOARCH=amd64 go build cw.go
	shasum -a 256 cw

build-linux: ## Build for Linux OS
	GOOS=linux GOARCH=amd64 go build cw.go
	shasum -a 256 cw

build-windows: ## Build for Windows OS
	GOOS=windows GOARCH=amd64 go build cw.go
	shasum -a 256 cw.exe

exec: ## Exec cw on a Symfony 5 project
	$(eval path ?= ../strangebuzz.com)
	./cw $(path)

clean: cw ## Clean the current executable
	rm ./cw

clean-all: cw cw.exe ## Clean all executable
	rm ./cw ./cw.exe

deps: clean ## Clean deps
	go get -d -v ./...

## —— Tests ✅ —————————————————————————————————————————————————————————————————
test: ## Run all tests
	go test -count=1 -v ./...

## —— Coding standards ✨ ——————————————————————————————————————————————————————
lint: ## Run gofmt simplify and lint
	gofmt -s -l -w .
