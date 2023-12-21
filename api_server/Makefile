.PHONY: test
test: ## Run all tests of go.
	make test-unit
	make test-repository
	make test-integration

.PHONY: fmt
fmt: ## Run fmt.
	make lint-fix
	make update-mock

.PHONY: lint-fix
lint-fix: ## Run golangci-lint fix.
	golangci-lint run ./... -v --fix

.PHONY: fmttf
fmttf: ## Run terraform fmt.
	terraform fmt -recursive

.PHONY: check-mock
check-mock: ## Check mock is up-to-date
	gomockhandler -config=./gomockhandler.json check

.PHONY: update-mock
update-mock: ## Update mock
	gomockhandler -config=./gomockhandler.json mockgen