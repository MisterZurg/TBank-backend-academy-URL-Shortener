.PHONY: all fmt tidy lint test
all: fmt tidy lint test

fmt:
	go fmt ./...

tidy:
	go mod tidy -v

lint:
	golangci-lint run

test:
	go clean -testcache
	go test -v ./...

# Docker CumPose Version
.PHONY: up down
up:
	docker compose -f ./docker-compose.yml rm && \
	docker compose -f ./docker-compose.yml build --no-cache && \
	docker compose -f ./docker-compose.yml up

down:
	docker-compose -f ./docker-compose.yml down


# k8s version
.PHONY: cluster deploy destroy
cluster:
	kind create cluster

destroy:
	kubectl delete statefulsets/clickhouse --namespace development