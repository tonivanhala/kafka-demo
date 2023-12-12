package main

import (
	"context"
	"log"
	"os"

	"github.com/IBM/sarama"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <topic name>", os.Args[0])
	}
	topicName := os.Args[1]

	brokers := []string{"localhost:9093"}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	group, err := sarama.NewConsumerGroup(brokers, topicName, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer group.Close()

	ctx := context.Background()
	for {
		handler := ConsumerGroupHandler{}
		err := group.Consume(ctx, []string{topicName}, handler)
		if err != nil {
			log.Fatalf("Error consuming partition: %v", err)
		}
	}
}

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Message received:\n\t%s\nOffset:\n\t%d", msg.Value, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}
