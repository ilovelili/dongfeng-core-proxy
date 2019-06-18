#!/bin/bash
set -e

cd ..

# docker build
docker build -t ilovelili/dongfeng-core-proxy .

echo "Bye"