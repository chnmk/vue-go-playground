package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"

	reader "github.com/chnmk/vue-go-playground/main/db/sqlite"
	transport "github.com/chnmk/vue-go-playground/main/transport/rest"

	_ "modernc.org/sqlite"
)

func main() {
	// Initialize databases
	sqlite := reader.Init()
	defer sqlite.Close()
	transport.SQLiteDB = sqlite

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

	// Keep the server alive until exit
	runtime.Goexit()
	fmt.Println("Exit")
}
