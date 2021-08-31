.PHONY: run
run: ## Run the server
	go run cmd/server/main.go

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

.PHONY: build
build: ## Compile the golang application
	CGO_ENABLED=0 go build -o bin/server cmd/server/main.go
	# CGO_ENABLED=0 go build -o bin/s3_sqs_worker cmd/s3_sqs_worker/main.go
	# CGO_ENABLED=0 go build -o bin/migrate cmd/migrate/main.go

lint:
	golangci-lint --color 'auto' run ./...

.PHONY: help
help:
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
