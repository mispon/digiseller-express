#!/bin/bash

# 1. Install docker and compose

### Docker and docker compose prerequisites
sudo apt-get install curl
sudo apt-get install gnupg
sudo apt-get install ca-certificates
sudo apt-get install lsb-release

###
sudo apt-get remove docker docker-engine docker.io containerd runc

### Update the apt package index and install packages to allow apt to use a repository over HTTPS:
sudo apt-get update

sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release -y

### Add Docker’s official GPG key:
sudo mkdir -m 0755 -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

### Use the following command to set up the repository:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

### Update the apt package index:
sudo apt-get update -y

### Install Docker Engine, containerd, and Docker Compose.
echo ""
echo "--------------------------------"
echo "Install docker"
echo "--------------------------------"

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y

### Post install
sudo groupadd docker || exit
sudo usermod -aG docker "$USER"
newgrp docker


# 2. Open ports
echo ""
echo "--------------------------------"
echo "Open ports for app and db admin"
echo "--------------------------------"

### Install ufw
apt install ufw -y

### Open ports for an app and adminer
sudo ufw allow 8080
sudo ufw allow 8082


# 3. Input envs
echo ""
echo "--------------------------------"
echo "Setup seller's data"
echo "--------------------------------"

echo "Enter SELLER_ID:"
read -r seller_id

echo "Enter SELLER_API_KEY:"
read -r seller_api_key

echo "Enter PG_USER:"
read -r pg_username

echo "Enter PG_PASS:"
read -r pg_password

echo "Enter TG_USER:"
read -r tg_username

echo "SELLER_ID=$seller_id
SELLER_API_KEY=$seller_api_key
PG_USER=$pg_username
PG_PASS=$pg_password
TG_USER=$tg_username" > .env


# 4. Run the app
echo ""
echo "--------------------------------"
echo "Start digi-express app"
echo "--------------------------------"

docker compose up -d