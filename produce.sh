#!/usr/bin/env bash

# Produce messages to a topic named with the first argument

if [ $# -eq 0 ]; then
  echo "Usage: produce.sh <topic-name>"
  exit 1
fi

docker run -it --rm --network kafka \
  bitnami/kafka kafka-console-producer.sh --broker-list kafka:9092 \
  --topic $1
