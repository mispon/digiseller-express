#!/bin/bash

# Update system
sudo apt update -y && sudo apt upgrade -y

# Install git if not exist
sudo apt install git

# Create service folder
mkdir digi-express && cd digi-express

# Download service files
git clone https://gist.github.com/8613e6a133d2eab625c60ffcf70c9e9c.git .

# Grant access and run service
chmod 777 repair.sh
chmod 777 update.sh
chmod 777 run.sh
sudo ./run.sh