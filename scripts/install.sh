#!/bin/bash

# Update system
sudo apt update -y && sudo apt upgrade -y

# Create service folder
mkdir digi-express && cd digi-express

# Download service files
curl -O "https://raw.githubusercontent.com/mispon/digiseller-express/master/docker-compose.yaml"
curl -O "https://raw.githubusercontent.com/mispon/digiseller-express/master/create_tables.sql"
curl -O "https://raw.githubusercontent.com/mispon/digiseller-express/master/scripts/run.sh"
curl -O "https://raw.githubusercontent.com/mispon/digiseller-express/master/scripts/update.sh"
curl -O "https://raw.githubusercontent.com/mispon/digiseller-express/master/scripts/repair.sh"

# Grant access and run service
chmod 777 repair.sh
chmod 777 update.sh
chmod 777 run.sh
sudo ./run.sh