#!/usr/bin/env bash

docker-compose -f docker-compose.dev.yml up -d

docker attach g3_server_dev