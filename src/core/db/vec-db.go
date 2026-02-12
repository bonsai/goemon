package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var GlobalVectorDB *VectorDB

type VectorDB struct {
	db *sql.DB
}

func InitGlobalVectorDB() error {
	db, err := sql.Open("sqlite3", "vec.db")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS vector_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		embedding BLOB,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	GlobalVectorDB = &VectorDB{db: db}
	return nil
}

func (vdb *VectorDB) GetRandomFragments(n int) ([]string, error) {
	query := `SELECT content FROM vector_records ORDER BY RANDOM() LIMIT ?`
	rows, err := vdb.db.Query(query, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fragments []string
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return nil, err
		}
		fragments = append(fragments, content)
	}
	return fragments, nil
}

func (vdb *VectorDB) Close() {
	if vdb.db != nil {
		vdb.db.Close()
	}
}
