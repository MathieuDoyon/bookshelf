#!make

.PHONY: help

include .env
export $(shell sed 's/=.*//' .env)

# TODO, I have to fix this command since we include .env it change the return of 
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# install: package.json ## install dependencies
# 	@yarn
install: ## install dependencies
	@dep ensure

# setup-env:
# 	@echo $$GO_PROJECT_DIR

serve: serve-dev

serve-dev: ## serve go with hot reload
	@docker-compose up
