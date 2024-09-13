package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type vueMessage struct {
	Message string `json:"message"`
}

func RestHandler(w http.ResponseWriter, r *http.Request) {
	var decoded vueMessage
	fmt.Println("New request to REST")

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Got the following message through REST: %s\n", decoded.Message)
	w.Write([]byte("REST API request success"))
}
