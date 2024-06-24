lint: ## Lint the source code
	@echo "🧹 Cleaning go.mod.."
	@go mod tidy
	@echo "🧹 Formatting files.."
	@go fmt ./...
	@echo "🧹 Vetting go.mod.."
	@go vet ./...

test: ## Run tests
	@echo "⏱️ Running tests.."
	@go test -v -count=1 -race -shuffle=on -coverprofile=coverage.txt ./...
