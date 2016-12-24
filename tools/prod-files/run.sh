#!/usr/bin/env bash

docker load -i vhustle.img
docker load -i nginxvhustle.img

tar -jxvf frontend.tar.bz2

docker-compose -p vhustle -f prod.docker-compose.yml up -d
