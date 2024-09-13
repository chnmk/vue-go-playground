package producer

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce() {
	topic := "playground"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("Hello from vue-go-playground through Kafka! No. 1")},
		kafka.Message{Value: []byte("Hello from vue-go-playground through Kafka! No. 2")},
		kafka.Message{Value: []byte("Hello from vue-go-playground through Kafka! No. 3")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
