# Very Simple Kafka

## Run Kafka and Zookeper

```bash
docker-compose up -d
```

Note that Zookeeper is not required for recent versions of Kafka.

## Create a topic, produce and consume messages

```bash
./create-topic.sh topic-name
./produce.sh topic-name
./consume.sh topic-name
```

## Consume using Go client

```bash
cd consumer
go run main.go topic-name
```

## Consume using Go client with consumer groups

```bash
cd consumer-group
go run main.go topic-name
```
