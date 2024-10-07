package transport

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	reader "github.com/chnmk/vue-go-playground/main/db/sqlite"
)

var SQLiteDB *sql.DB

func SqliteHandler(w http.ResponseWriter, r *http.Request) {
	var decoded vueMessage

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := reader.Read(SQLiteDB, decoded.Message)

	result := fmt.Sprintf("Added the folllowing text: %s, last insert id: %d", decoded.Message, id)
	w.Write([]byte(result))
}
