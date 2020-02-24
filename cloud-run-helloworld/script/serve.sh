#!/bin/bash

cd $(dirname $0)/../cmd/local
skaffold dev --no-prune=false --no-prune-children=false --cache-artifacts=false
