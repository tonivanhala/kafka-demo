#!/usr/bin/env bash

# Consume messages from a topic named with the first argument

if [ $# -eq 0 ]; then
  echo "Usage: consume.sh <topic-name>"
  exit 1
fi

docker run -it --rm --network kafka \
  bitnami/kafka kafka-console-consumer.sh --bootstrap-server kafka:9092 \
  --topic $1 --from-beginning
