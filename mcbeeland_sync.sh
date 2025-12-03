#!/bin/sh

docker build --network=host . -t svc-discord:latest
docker save svc-discord:latest | zstd -2 - | pv -b | ssh mikinol-serv "zstd -d - | docker load"
