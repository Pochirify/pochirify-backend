POCHIRIFY_BACKEND_IMAGE_TAG := asia-northeast1-docker.pkg.dev/pochirify-dev/pochirify-backend/pochirify-backend:latest

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
		--allow-unauthenticated
