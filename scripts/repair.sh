#!/bin/bash

echo ""
echo "--------------------------------"
echo "Stop and remove old setup"
echo "--------------------------------"
sudo docker compose stop && sudo docker compose rm -f
sudo docker volume rm digi-express_dbdata

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

echo ""
echo "--------------------------------"
echo "Start digi-express app"
echo "--------------------------------"
sudo docker compose up -d