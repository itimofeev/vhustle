#!/usr/bin/env bash

rsync -a /Users/ilyatimofee/prog/js/hustlesa-ui/sync/ ilyaufo@188.166.26.165:/home/ilyaufo/vhustle/frontend/
scp -r target/* ilyaufo@188.166.26.165:/home/ilyaufo/vhustle
