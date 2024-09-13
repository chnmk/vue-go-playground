package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"

	// transport_grpc "github.com/chnmk/vue-go-playground/main/transport/grpc/server"

	transport "github.com/chnmk/vue-go-playground/main/transport/rest"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Read SQL
	// reader.ReadSQLite()

	// Connect to Redis
	// go redis_client.ExampleClient()

	// Connect to Kafka
	// go producer.Produce()

	// GraphQL
	// go transport_grqphql.HelloFromGraphQL()

	// grpc server
	// go transport_grpc.Serve()

	// Run handlers
	http.HandleFunc("/api/REST", transport.RestHandler)
	http.HandleFunc("/api/gRPC", transport.GrpcHandler)
	http.HandleFunc("/api/GraphQL", transport.GraphqlHandler)
	http.HandleFunc("/api/Kafka", transport.KafkaHandler)
	http.HandleFunc("/api/RabbitMQ", transport.RabbitmqHandler)
	http.HandleFunc("/api/SQLite", transport.SqliteHandler)
	http.HandleFunc("/api/PostgreSQL", transport.PostgreHandler)
	http.HandleFunc("/api/Redis", transport.RedisHandler)
	http.HandleFunc("/api/Elastic", transport.ElasticHandler)
	http.HandleFunc("/api/Clickhouse", transport.ClickhouseHandler)
	http.HandleFunc("/api/R", transport.RlangHandler)

	// Serve frontend app
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Create listener
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Panic(err)
	}

	// Serve in a goroutine
	go func() {
		log.Panic(
			http.Serve(ln, nil),
		)
	}()

	// Keep server goroutine alive until exit
	runtime.Goexit()
	fmt.Println("Exit")
}
