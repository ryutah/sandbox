#!/bin/sh

set -eux

cd $(dirname $0)/../app

gcloud builds submit -t gcr.io/$(gcloud config get-value project 2>/dev/null)/helloworld:latest .
