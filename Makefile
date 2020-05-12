## —— 🐝 The Cache Watcher Makefile 🐝 —————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Project 🐝 ———————————————————————————————————————————————————————————————
run: ## Run the main go file on a Symfony 5 project
	go run cw.go ../strangebuzz.com

build: ## Build the sfcw executable for the current platform
	go build cw.go
	shasum -a 256 cw
	# To build the Linux executable run
	# env GOOS=darwin GOARCH=amd64 go build cw.go
	# env GOOS=linux GOARCH=amd64 go build cw.go
	# env GOOS=windows GOARCH=amd64 go build cw.go

exec: ## Run cc on a Symfony 5 project
	./cc ../strangebuzz.com

clean: sfcw ## Clean the current executable
	rm ./cw

deps: clean ## Clean deps
	go get -d -v ./...

## —— Tests ✅ —————————————————————————————————————————————————————————————————
test: ## Run all tests
	go test -count=1 -v ./...

## —— Coding standards ✨ ——————————————————————————————————————————————————————
lint: ## Run gofmt simplify and lint
	gofmt -s -l -w .
