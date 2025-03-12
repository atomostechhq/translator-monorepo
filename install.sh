#!/bin/bash

echo "============================"
echo " Checking Docker installation"
echo "============================"

if ! command -v docker &> /dev/null
then
    echo "Docker not found! Installing Docker..."
    curl -fsSL https://get.docker.com -o get-docker.sh
    sh get-docker.sh
else
    echo "Docker is installed."
fi

echo "============================"
echo " Checking Docker Compose"
echo "============================"

if ! command -v docker compose &> /dev/null
then
    echo "Docker Compose not found! Installing..."
    sudo curl -L \"https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)\" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
else
    echo "Docker Compose is installed."
fi

echo "============================"
echo " Pulling images & starting containers"
echo "============================"
docker compose pull
docker compose up -d
