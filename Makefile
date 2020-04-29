## —— 🐝 The SymFony Cache Watcher Makefile 🐝 —————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Project 🐝 ———————————————————————————————————————————————————————————————
run: ## Run the main go file on a Symfony 5 project
	go run sfcw.go ../strangebuzz.com

build: ## Build the sfcw executable
	go build sfcw.go

exec: ## Run cc on a Symfony 5 project
	./cc ../strangebuzz.com

clean: sfcw ## Clean the current executable
	rm ./sfcw

deps: clean ## Clean deps
	go get -d -v ./...

## —— Tests ✅ —————————————————————————————————————————————————————————————————
test: ## Run all tests
	go test

## —— Coding standards ✨ ——————————————————————————————————————————————————————
lint: ## Run gofmt simplify and lint
	gofmt -s -l -w .
