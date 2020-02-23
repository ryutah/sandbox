#!/bin/bash

cd $(dirname $0)/../resources

set -eux

gcloud deployment-manager deployments update gke-google-managed-ssl --config cluster.yaml

