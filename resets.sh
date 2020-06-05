#!/bin/zsh
clear
docker container ls -a
docker ps
docker container stop mongodb-server
docker container stop mongo-client
docker container prune -f
docker network prune -f
docker volume prune -f
docker images prune
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)
docker container ls -a
docker ps