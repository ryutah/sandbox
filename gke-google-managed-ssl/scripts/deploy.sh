#!/bin/sh

cd $(dirname $0)/..

set -eux

gcloud container clusters get-credentials example-cluster --zone asia-northeast1-a
kustomize build k8s/overlays/production | kubectl apply -f -
