#!/usr/bin/env bash

docker-compose -p nginx -f docker-compose.yml up -d --build --remove-orphans
