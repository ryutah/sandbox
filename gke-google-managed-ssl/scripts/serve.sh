#!/bin/bash

cd $(dirname $0)/..

set +C -eux

kustomize build k8s/base > k8s.yaml
skaffold dev --no-prune=false --cache-artifacts=false --no-prune-children=false
