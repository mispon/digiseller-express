#!/bin/bash

# Download service archive
curl -O "url"

# Update system packages
sudo apt update -y && apt upgrade -y

# Unzip service archive
apt install unzip -y
unzip digi-express.zip
cd digi-express || exit

# Grant access and run service
chmod 777 update.sh
chmod 777 run.sh
sudo ./run.sh