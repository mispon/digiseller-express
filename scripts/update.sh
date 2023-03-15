# Stop the app and remove old containers
docker compose stop && docker compose rm -f

# Download new app's image and run
docker compose up -d