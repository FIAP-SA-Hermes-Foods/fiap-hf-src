#!/bin/bash

OUT_IP_ADDR=$(docker network ls | grep hermes_foods_net_dev)

if [ ! -z "$OUT_IP_ADDR" -a "$OUT_IP_ADDR" != " " ]; then
    echo "network already exists, initializing..."
else
    docker network create --subnet 204.7.9.0/24 hermes_foods_net_dev
fi
