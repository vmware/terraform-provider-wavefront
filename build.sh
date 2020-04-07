#!/bin/bash

VERSION=$(cat version)
echo "building terraform-provider-wavefront_${VERSION}-linux-amd64..."
env GOOS=linux GOARCH=amd64 go build -o terraform-provider-wavefront_${VERSION}-linux-amd64
echo "building terraform-provider-wavefront_${VERSION}-darwin-amd64..."
env GOOS=darwin GOARCH=amd64 go build -o terraform-provider-wavefront_${VERSION}-darwin-amd64
