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
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y


# 2. Open ports

### Install ufw
apt install ufw -y

### Open ports for an app and adminer
sudo ufw allow 8080
sudo ufw allow 8082


# 3. Run app

docker compose up -d