.PHONY: run
run: ## Run the server
	go run cmd/server/main.go

.PHONY: worker
worker:
	go run cmd/worker/main.go --worker_function=get_html_worker

.PHONY: migrate
migrate: migrate_dev migrate_test

.PHONY: migrate_dev
migrate_dev: ## Migrate the dev database
	go run cmd/migrate/main.go --database_name=keyword_recommendation_development

.PHONY: migrate_test
migrate_test: ## Migrate the test database
	go run cmd/migrate/main.go --database_name=keyword_recommendation_test

.PHONY: tools
tools: ## Installs dev tools
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go get %

.PHONY: watch
watch: ## Start the server and watch for changes
	reflex -r '\.go' -s go run cmd/server/main.go

.PHONY: gen
gen: ## Generate golang code
	go generate ./...

.PHONY: test
test: migrate_test ## Run the app tests
	go test -v ./...

.PHONY: test_color
testc: ## Run the app tests with color output
	gotest -v ./...

.PHONY: image
image: ## Build golang image
	docker build . -t crawler_api:0.0.4
	docker tag crawler_api:0.0.4 hub.normans.co.za/crawler:0.0.4
	docker push hub.normans.co.za/crawler:0.0.4

.PHONY: build
build: ## Compile the golang application
	CGO_ENABLED=0 go build -o bin/server cmd/server/main.go
	CGO_ENABLED=0 go build -o bin/worker cmd/worker/main.go
	# CGO_ENABLED=0 go build -o bin/migrate cmd/migrate/main.go

lint:
	golangci-lint --color 'auto' run ./...

.PHONY: help
help:
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
