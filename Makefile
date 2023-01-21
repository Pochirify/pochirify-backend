# OS   := $(shell uname | awk '{print tolower($$0)}')
# ARCH := $(shell case $$(uname -m) in (x86_64) echo amd64 ;; (aarch64) echo arm64 ;; (*) echo $$(uname -m) ;; esac)
include .env
POCHIRIFY_BACKEND_IMAGE_TAG := asia-northeast1-docker.pkg.dev/pochirify-dev/pochirify-backend/pochirify-backend:latest

DEV_DIR   := $(shell pwd)/dev
BIN_DIR   := $(DEV_DIR)/bin
TOOLS_DIR := $(DEV_DIR)/tools
TOOLS_SUM := $(TOOLS_DIR)/go.sum
YQ_VERSION := 4.14.1
YO := $(abspath $(BIN_DIR)/yo)
BUILD_TOOLS := cd $(TOOLS_DIR) && go build -o

.PHONY: yo
yo: $(YO)
$(YO): $(TOOLS_SUM)
	@$(BUILD_TOOLS) $(YO) go.mercari.io/yo

# login: https://cloud.google.com/spanner/docs/getting-started/set-up?hl=ja
pochirify-yo: $(YO)
	@$(YO) \
		pochirify-dev \
		pochirify \
		pochirify-server \
		--out ./internal/handler/db/spanner/yo

gqlgen:
	go run github.com/99designs/gqlgen generate

up:
	docker compose up

down:
	docker compose down

# とりあえず手動でデプロイできるようにした
.PHONY: backend-image
backend-image:
	sudo docker buildx build --platform linux/amd64 -t $(POCHIRIFY_BACKEND_IMAGE_TAG) -f ./Dockerfile .

.PHONY: backend-registry
backend-registry:
	docker push $(POCHIRIFY_BACKEND_IMAGE_TAG)

.PHONY: backend-deploy
backend-deploy:
	gcloud beta run deploy "pochirify-backend" \
		--project pochirify-dev \
		--image $(POCHIRIFY_BACKEND_IMAGE_TAG) \
		--platform managed \
		--region asia-northeast1 \
		--allow-unauthenticated \
		--update-env-vars GCP_PROJECT_ID=$(GCP_PROJECT_ID) \
		--update-env-vars SPANNER_INSTANCE_ID=$(SPANNER_INSTANCE_ID) \
		--update-env-vars SPANNER_DATABASE_ID=$(SPANNER_DATABASE_ID)
