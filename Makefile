SERVICE := apigw

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
	go install github.com/bufbuild/protoc-gen-validate@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install github.com/lyft/protoc-gen-star/protoc-gen-debug@latest

.PHONY: buf
buf:
	buf generate

local:
	go run main.go

deploy: ## Deploy
	fly deploy -a ${SERVICE} --no-cache

conf secret: ## Upddate config
	fly secrets import -a ${SERVICE} < fly.env

bash: ## Enter bash
	fly ssh console -a ${SERVICE}

docker:
	docker compose build ${SERVICE}
	docker compose up ${SERVICE}

updb:
	docker compose up -d ${SERVICE}-db --force-recreate

db:
	docker compose exec ${SERVICE}-db bash