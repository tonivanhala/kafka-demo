#!/usr/bin/env bash

# Create a topic named with the first argument

if [ $# -eq 0 ]; then
  echo "Usage: create-topic.sh <topic-name>"
  exit 1
fi

docker run -it --rm --network kafka \
  bitnami/kafka kafka-topics.sh --create --bootstrap-server kafka:9092 \
  --replication-factor 1 --partitions 1 --topic $1
