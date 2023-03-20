#!/bin/bash

# Update system
sudo apt update -y && sudo apt upgrade -y

# Create service folder
mkdir digi-express && cd digi-express

# Download service files
curl -O "https://gist.githubusercontent.com/mispon/8613e6a133d2eab625c60ffcf70c9e9c/raw/4893d4b3dee8f6e668abe56867ddd466a73d5adf/docker-compose.yaml"
curl -O "https://gist.githubusercontent.com/mispon/8613e6a133d2eab625c60ffcf70c9e9c/raw/4893d4b3dee8f6e668abe56867ddd466a73d5adf/create_tables.sql"
curl -O "https://gist.githubusercontent.com/mispon/8613e6a133d2eab625c60ffcf70c9e9c/raw/4893d4b3dee8f6e668abe56867ddd466a73d5adf/run.sh"
curl -O "https://gist.githubusercontent.com/mispon/8613e6a133d2eab625c60ffcf70c9e9c/raw/4893d4b3dee8f6e668abe56867ddd466a73d5adf/update.sh"
curl -O "https://gist.githubusercontent.com/mispon/8613e6a133d2eab625c60ffcf70c9e9c/raw/4893d4b3dee8f6e668abe56867ddd466a73d5adf/repair.sh"

# Grant access and run service
chmod 777 repair.sh
chmod 777 update.sh
chmod 777 run.sh
sudo ./run.sh