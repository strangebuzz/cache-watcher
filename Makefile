## —— 🐝 The Strangebuzz cc Makefile 🐝 ————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Project 🐝 ———————————————————————————————————————————————————————————————
run: ## Run the main go file
	go run cc.go ../strangebuzz.com

build: ## Build the cc executable
	go build cc.go

exec: ## Run cc with a Symfony 5 project
	./cc ../strangebuzz.com

clean: cc ## Clean the current executable
	rm ./cc

deps: clean ## Clean deps
	go get -d -v ./...

update-deps: ## Update all deps
	go get -u; \
	go mod tidy; \

## —— Tests ✅ —————————————————————————————————————————————————————————————————
test: ## Run all tests
	go test

## —— Coding standards ✨ ——————————————————————————————————————————————————————
lint: ## Run gofmt simplify and lint
	gofmt -s -l -w .
