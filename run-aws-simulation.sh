#!/usr/bin/env bash
# We need to set the OS to linux.
# AWS Lambda runs in a linux container, which we're not doing on my MacBook.
set -e
GOOS=linux go build -tags aws -o ball-clock-simulator-lambda.out
sam local start-api

