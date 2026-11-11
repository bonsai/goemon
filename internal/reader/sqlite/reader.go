package sqlite

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"goemon/internal/document"
	_ "github.com/mattn/go-sqlite3"
)

// Reader opens an existing SQLite database in read-only mode.
type Reader struct { Path string; Table string; Limit int }

func (r Reader) ReadAll(ctx context.Context) ([]document.Document, error) {
	abs, err := filepath.Abs(r.Path); if err != nil { return nil, err }
	u := url.URL{Scheme:"file", Path:abs}
	q := u.Query(); q.Set("mode","ro"); q.Set("_query_only","1"); u.RawQuery=q.Encode()
	db, err := sql.Open("sqlite3", u.String()); if err != nil { return nil, err }
	defer db.Close()
	if err := db.PingContext(ctx); err != nil { return nil, err }
	tables, err := db.QueryContext(ctx, "SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%' ORDER BY name")
	if err != nil { return nil, err }
	var names []string
	for tables.Next() { var n string; if err:=tables.Scan(&n); err!=nil { tables.Close(); return nil,err }; names=append(names,n) }
	if err:=tables.Err(); err!=nil { tables.Close(); return nil,err }; tables.Close()
	out:=make([]document.Document,0)
	for _, table := range names {
		if r.Table!="" && r.Table!=table { continue }
		// Table names originate from sqlite_master; quote identifiers defensively.
		quoted := "\"" + strings.ReplaceAll(table, "\"", "\"\"") + "\""
		limit:=r.Limit; if limit<=0 { limit=1000 }
		rows,err:=db.QueryContext(ctx,fmt.Sprintf("SELECT * FROM %s LIMIT ?",quoted),limit); if err!=nil{return nil,err}
		cols,err:=rows.Columns(); if err!=nil{rows.Close();return nil,err}
		for rows.Next() {
			values:=make([]any,len(cols)); ptrs:=make([]any,len(cols))
			for i:=range values { ptrs[i]=&values[i] }
			if err:=rows.Scan(ptrs...); err!=nil { rows.Close(); return nil,err }
			record:=map[string]any{}
			for i,c:=range cols {
				switch v:=values[i].(type) { case []byte: record[c]=string(v); default: record[c]=v }
			}
			raw,err:=json.Marshal(record); if err!=nil{rows.Close();return nil,err}
			sum:=sha256.Sum256(raw)
			out=append(out,document.Document{ID:fmt.Sprintf("sqlite:%s:%s",table,hex.EncodeToString(sum[:])),Source:"sqlite",SourceID:table,Title:table,Body:string(raw),ContentType:"application/json",Metadata:map[string]any{"table":table,"record":record},RawReference:r.Path})
		}
		if err:=rows.Err(); err!=nil{rows.Close();return nil,err}; rows.Close()
	}
	return out,nil
}
