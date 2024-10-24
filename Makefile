SERVICE := noah

all: help

help: ## Show help messages
	@echo "Container - ${SERVICE} "
	@echo
	@echo "Usage:\tmake COMMAND"
	@echo
	@echo "Commands:"
	@sed -n '/##/s/\(.*\):.*##/  \1#/p' ${MAKEFILE_LIST} | grep -v "MAKEFILE_LIST" | column -t -c 2 -s '#'

.PHONY: init
init:  ## Init project
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/PaulXu-cn/go-mod-graph-chart/gmchart@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/lyft/protoc-gen-star/protoc-gen-debug@latest

.PHONY: buf
buf:
	buf generate

local: ## Run service locally 
	go run main.go run -c config.local.yaml

build: ## Build or rebuild docker image
	#docker compose build --no-cache
	docker compose build ${SERVICE}

up:
	docker compose up -d ${SERVICE} --force-recreate

start:
	docker compose start ${SERVICE}

stop:
	docker compose stop ${SERVICE}

bash: ## Enter bash
	docker compose exec ${SERVICE} bash

logs:
	docker compose logs -f ${SERVICE}

dbup:
	docker compose up -d ${SERVICE}-db --force-recreate

dbrm dbclean:
	docker compose rm -f ${SERVICE}-db
	#docker volume rm ${SERVICE}-db

dbbash:
	docker compose exec ${SERVICE}-db bash
