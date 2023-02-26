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
GQLGENC := $(abspath $(BIN_DIR)/gqlgenc)
GQLGEN := $(abspath $(BIN_DIR)/gqlgen)
BUILD_TOOLS := cd $(TOOLS_DIR) && go build -o

.PHONY: yo
yo: $(YO)
$(YO): $(TOOLS_SUM)
	@$(BUILD_TOOLS) $(YO) go.mercari.io/yo

.PHONY: gqlgenc
gqlgenc: $(GQLGENC)
$(GQLGENC): $(TOOLS_SUM)
	@$(BUILD_TOOLS) $(GQLGENC) github.com/Yamashou/gqlgenc

.PHONY: gqlgen
gqlgen: $(GQLGEN)
$(GQLGEN): $(TOOLS_SUM)
	@$(BUILD_TOOLS) $(GQLGEN) github.com/99designs/gqlgen

# login: https://cloud.google.com/spanner/docs/getting-started/set-up?hl=ja
pochirify-yo: $(YO)
	@$(YO) \
		pochirify-dev \
		pochirify \
		pochirify-server \
		--out ./internal/handler/db/spanner/yo

pochirify-gqlgenc: $(GQLGENC)
	@$(GQLGENC)

# TODO: --configdir効かせる
# shopify-gqlgenc: $(GQLGENC)
# @$(GQLGENC) generate --configdir ./e2etests/shopify

pochirify-gqlgen: $(GQLGEN)
	@$(GQLGEN)
# gqlgen:
# 	go run github.com/99designs/gqlgen generate

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
		--update-env-vars SPANNER_DATABASE_ID=$(SPANNER_DATABASE_ID) \
		--update-env-vars FINCODE_API_KEY=$(FINCODE_API_KEY) \
		--update-env-vars FINCODE_BASE_URL=$(FINCODE_BASE_URL) \
		--update-env-vars IS_PAYPAY_PRODUCTION=$(IS_PAYPAY_PRODUCTION) \
		--update-env-vars PAYPAY_API_KEY_ID=$(PAYPAY_API_KEY_ID) \
		--update-env-vars PAYPAY_API_SECRET=$(PAYPAY_API_SECRET) \
		--update-env-vars PAYPAY_MERCHANT_ID=$(PAYPAY_MERCHANT_ID) \
		--update-env-vars SHOPIFY_ADMIN_ACCESS_TOKEN=$(SHOPIFY_ADMIN_ACCESS_TOKEN)

backend-deploy-all: backend-image backend-registry backend-deploy

.PHONY: e2etest
e2etest:
	go test -tags=e2e -shuffle=on ./e2etests/...
