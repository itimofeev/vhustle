#!/usr/bin/env bash

export GOPATH=/Users/ilyatimofee/prog/axxonsoft/

PROJECT_PATH=/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/vhustle
FRONTEND_PROJECT_PATH=/Users/ilyatimofee/prog/js/hustlesa-ui

rm ${PROJECT_PATH}/target/*
mkdir ${PROJECT_PATH}/target


export GOOS=linux
export GOARCH=amd64
go build -v github.com/itimofeev/vhustle/main/vhustle


docker build --force-rm=true -t vhustle -f ${PROJECT_PATH}/tools/vhustle.Dockerfile .
docker save -o "$PROJECT_PATH/target/vhustle.img" "vhustle"


rm vhustle

cp ${PROJECT_PATH}/tools/prod-files/* ${PROJECT_PATH}/target/

echo 'building frontend'

#npm build ${FRONTEND_PROJECT_PATH}
cp -r ${FRONTEND_PROJECT_PATH}/build ${PROJECT_PATH}/target/frontend
cd target
tar -jcvf ${PROJECT_PATH}/target/frontend.tar.bz2 frontend

rm -r frontend
