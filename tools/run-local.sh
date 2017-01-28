#!/usr/bin/env bash

TOOLS_PATH=/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/vhustle/tools
TARGET_PATH=/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/vhustle/target

${TOOLS_PATH}/build.sh

docker load -i ${TARGET_PATH}/vhustle.img
docker-compose -p vhustle -f ${TOOLS_PATH}/dev.docker-compose.yml up -d
