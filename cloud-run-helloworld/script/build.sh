#!/bin/bash

cd $(dirname $0)/..

set -eux

gcloud builds submit -t gcr.io/$(gcloud config get-value project 2>/dev/null)/cloud-run-started .
