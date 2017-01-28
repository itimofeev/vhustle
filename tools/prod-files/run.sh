#!/usr/bin/env bash

docker load -i vhustle.img

tar -jxvf frontend.tar.bz2
chown ilyaufo frontend

docker-compose -p vhustle -f prod.docker-compose.yml up -d --build
