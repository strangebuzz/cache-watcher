# Setup ————————————————————————————————————————————————————————————————————————
SHELL         = bash
PROJECT       = cc
#.PHONY       = # Not needed for now

## —— 🐝 The Strangebuzz cc Makefile 🐝 ————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Project 🐝 ———————————————————————————————————————————————————————————————
run: ## Run the main go file
	go run hello-world.go

build: ## Build the cc executable
	go build hello-world.go

clean: hello-world ## Clean the current executable
	rm ./hello-world

## —— Tests ✅ —————————————————————————————————————————————————————————————————
#test: ## Launch main functional and unit tests
	#./vendor/bin/phpunit --testsuite=main --stop-on-failure

## —— Coding standards ✨ ——————————————————————————————————————————————————————
cs: ## Launch check style
	gofmt *.go

## —— Stats 📜 —————————————————————————————————————————————————————————————————
stats: ## Commits by the hour for the main author of this project
	$(GIT) log --author="$(GIT_AUTHOR)" --date=iso | perl -nalE 'if (/^Date:\s+[\d-]{10}\s(\d{2})/) { say $$1+0 }' | sort | uniq -c|perl -MList::Util=max -nalE '$$h{$$F[1]} = $$F[0]; }{ $$m = max values %h; foreach (0..23) { $$h{$$_} = 0 if not exists $$h{$$_} } foreach (sort {$$a <=> $$b } keys %h) { say sprintf "%02d - %4d %s", $$_, $$h{$$_}, "*"x ($$h{$$_} / $$m * 50); }'
