package transport

import (
	"fmt"
	"net/http"

	producer "github.com/chnmk/vue-go-playground/main/transport/kafka"
)

func KafkaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New request to Kafka")

	err := producer.Produce()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Kafka request success"))
}
