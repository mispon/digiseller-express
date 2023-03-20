#!/bin/bash

# Stop the app and remove old containers
sudo docker compose stop && sudo docker compose rm -f

# Download new app's image
sudo docker pull mispon/digi-express:latest

# And run
sudo docker compose up -d