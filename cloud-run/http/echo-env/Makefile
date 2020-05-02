.PHONY: all

CURDIR := $(shell pwd)

help: ## Print this help
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

DOCKER_IMAGE_VERSION = 1.0.0

set_project_id:
	$(eval project_id := $(shell gcloud config get-value project 2>/dev/null))

build_docker: set_project_id ## Build Docker Image
	gcloud --project ${project_id} --quiet builds submit -t gcr.io/${project_id}/echo-env:${DOCKER_IMAGE_VERSION}

deploy: set_project_id ## Deploy to Cloud Run
	gcloud --quiet run deploy echo-env \
		--image gcr.io/${project_id}/echo-env:${DOCKER_IMAGE_VERSION} \
		--platform managed \
		--region asia-northeast1 \
		--allow-unauthenticated
