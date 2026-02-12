package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var GlobalMailDB *sql.DB

func InitGlobalMailDB() error {
	var err error
	GlobalMailDB, err = sql.Open("sqlite3", "mail.db")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS mail_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender TEXT,
		recipient TEXT,
		subject TEXT,
		body TEXT,
		received_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = GlobalMailDB.Exec(query)
	return err
}
