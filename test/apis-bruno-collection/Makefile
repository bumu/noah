all: help

help: ## Show help messages
	@echo "Container - ${SERVICE} "
	@echo
	@echo "Usage:\tmake COMMAND"
	@echo
	@echo "Commands:"
	@sed -n '/##/s/\(.*\):.*##/  \1#/p' ${MAKEFILE_LIST} | grep -v "MAKEFILE_LIST" | column -t -c 2 -s '#'

install: ## Install bruno
	npm install -g @usebruno/cli

local: ## Run bruno in local
	bru run --env local

dev: ## Run bruno in dev
	bru run --env dev

test: ## Run bruno in test
	bru run --env test