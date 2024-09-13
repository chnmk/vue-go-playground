package producer

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce() error {
	topic := "playground"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Print(err)
		return errors.New("couldn't connect to Kafka")
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("Hello from vue-go-playground through Kafka! No. 1")},
		kafka.Message{Value: []byte("Hello from vue-go-playground through Kafka! No. 2")},
		kafka.Message{Value: []byte("Hello from vue-go-playground through Kafka! No. 3")},
	)
	if err != nil {
		log.Print(err)
		return errors.New("couldn't write messages to Kafka")
	}

	if err := conn.Close(); err != nil {
		log.Print(err)
		return errors.New("failed to close Kafka writer")
	}

	return nil
}
