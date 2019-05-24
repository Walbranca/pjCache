#!/usr/bin/env bash
rm -rf .temp
mkdir -p .temp/private-repositories/github.com/xMota14/
cp -r ${HOME}/go/src/github.com/xMota14/postgresAux/ .temp/private-repositories/github.com/xMota14/postgresAux/
sudo docker-compose build --no-cache pjcache
sh start.sh
