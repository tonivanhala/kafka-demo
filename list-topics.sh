#!/usr/bin/env bash

# List all topics

docker run -it --rm --network kafka \
  bitnami/kafka kafka-topics.sh --list --bootstrap-server kafka:9092