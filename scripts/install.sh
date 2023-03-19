#!/bin/bash

# Download service archive
curl -O "https://github.com/mispon/digiseller-express/releases/download/v1.0.0/digi-express.zip"

# Update system packages
sudo apt update -y
sudo apt upgrade -y

# Unzip service archive
sudo apt install unzip -y
unzip digi-express.zip
cd digi-express || exit

# Grant access and run service
chmod 777 update.sh
chmod 777 run.sh
sudo ./run.sh