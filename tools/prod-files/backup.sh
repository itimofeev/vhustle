#!/usr/bin/env bash

docker exec -t vhustle_db_1 pg_dumpall -c -U postgres > dump_`date +%d-%m-%Y"_"%H_%M_%S`.sql
docker exec -t vhustle_db_1 pg_dumpall -c -U postgres > dump_latest.sql
