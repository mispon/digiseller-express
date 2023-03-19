#!/bin/bash

# Stop the app and remove old containers
docker compose stop && docker compose rm -f

# Download new app's image
docker pull mispon/digi-express:latest

# And run
docker compose up -d