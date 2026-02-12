package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var GlobalDB *sql.DB

func InitGlobalDB() error {
	var err error
	GlobalDB, err = sql.Open("sqlite3", "imagestack.db")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS images (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT,
		prompt TEXT,
		model TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = GlobalDB.Exec(query)
	return err
}

func CloseDB() {
	if GlobalDB != nil {
		GlobalDB.Close()
	}
}
