package reader

import (
	"database/sql"
	"fmt"
	"os"
)

func Init() *sql.DB {
	_, errNoDB := os.Stat("./main/db/sqlite/sqlite.db")

	db, err := sql.Open("sqlite", "./main/db/sqlite/sqlite.db")
	if err != nil {
		fmt.Println(err)
	}

	if errNoDB == nil {
		return db
	}

	_, err = db.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY AUTOINCREMENT, text VARCHAR(255));")
	if err != nil {
		fmt.Println(err)
	}

	return db
}

func Read(db *sql.DB, str string) int64 {
	result, err := db.Exec("insert into test (text) values (:s)", sql.Named("s", str))
	if err != nil {
		fmt.Println(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	// rowsAffecred, _ := result.RowsAffected()

	fmt.Println("SQLite last insert ID: ", lastInsertId)
	// fmt.Println("SQLite rows affected: ", rowsAffecred)

	return lastInsertId
}
