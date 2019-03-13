#!/bin/bash
set -e
# move to root directory
cd ..
# docker build
docker build -t ilovelili/dongfeng-core-proxy . -f DockerFile
echo "Bye"