#!/bin/sh

set -eux

gcloud run deploy cloud-run-started \
  --platform managed \
  --port 8080 \
  --allow-unauthenticated \
  --region asia-northeast1 \
  --set-env-vars "GOOGLE_CLOUD_PROJECT=$(gcloud config get-value project 2>/dev/null)" \
  --image gcr.io/sandbox-hara/cloud-run-started
