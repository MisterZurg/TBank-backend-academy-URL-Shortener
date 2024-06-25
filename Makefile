########################################################################
### Before Push									      ##
########################################################################
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

########################################################################
### Docker CumPose Version										      ##
########################################################################
.PHONY: up down
up:
	docker compose -f ./docker-compose.yml rm && \
	docker compose -f ./docker-compose.yml build --no-cache && \
	docker compose -f ./docker-compose.yml up

down:
	docker-compose -f ./docker-compose.yml down


########################################################################
### K8S Version										      			  ##
########################################################################
.PHONY: cluster click nms redis traefik app prometheus grafana deploy destroy
cluster:
	kind create cluster --config ./k8s/kind-config.yml

deploy:	nms traefik click redis app prometheus grafana

nms:
	# namespace for traefik
	kubectl create namespace traefik
	# namespace storage chache db
	kubectl create namespace t-bank-storage
	# namespace for microservice
	kubectl create namespace t-bank-url-shortener
	# namespace for ,etrics
	kubectl create namespace t-bank-metrics

	kubectl get namespace

traefik:
	helm install traefik traefik/traefik -n traefik -f ./helm/traefik-values.yml

click:
	# config
	kubectl create -n storage -f ./k8s/tbank-ch/cm.yml
	# pvc
	kubectl create -n storage -f ./k8s/tbank-ch/pvc.yml
	# statefulset
	kubectl create -n storage -f ./k8s/tbank-ch/statefulset.yml
	# svc
	kubectl create -n storage -f ./k8s/tbank-ch/svc.yml

redis:
	# config
	kubectl create -n storage -f ./k8s/tbank-redis/cm.yml
	# pvc
	 kubectl create -n storage -f ./k8s/tbank-redis/pvc.yml
	# statefulset
	kubectl create -n storage -f ./k8s/tbank-redis/statefulset.yml
	# svc
	kubectl create -n storage -f ./k8s/tbank-redis/svc.yml


app:
	# config
	kubectl create -n api -f ./k8s/tbank-url-shortener/cm.yml
	# statefulset
	kubectl create -n api -f ./k8s/tbank-url-shortener/deploy.yml
	# svc
	kubectl create -n api -f ./k8s/tbank-url-shortener/svc.yml
	# traefick
	kubectl create -n api -f ./k8s/tbank-url-shortener/app-traefik-ingress.yml


prometheus:

grafana:

destroy:
	kubectl delete all --all -n t-bank-url-shortener
	kubectl delete all --all -n t-bank-storage
	kubectl delete all --all -n t-bank-metrics

	helm uninstall traefik traefik -n traefik
