package sqlite

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestReaderReadsExistingDatabaseWithoutChangingIt(t *testing.T) {
	path:=filepath.Join(t.TempDir(),"source.db")
	db,err:=sql.Open("sqlite3",path); if err!=nil { t.Fatal(err) }
	if _,err=db.Exec("CREATE TABLE notes (id INTEGER PRIMARY KEY, title TEXT); INSERT INTO notes(title) VALUES ('alpha'),('beta')"); err!=nil { t.Fatal(err) }
	if err=db.Close(); err!=nil { t.Fatal(err) }
	before,err:=os.ReadFile(path); if err!=nil { t.Fatal(err) }

	docs,err:=(Reader{Path:path,Table:"notes",Limit:10}).ReadAll(context.Background())
	if err!=nil { t.Fatal(err) }
	if len(docs)!=2 { t.Fatalf("want 2 records, got %d",len(docs)) }
	if docs[0].ID=="" || docs[0].ID==docs[1].ID { t.Fatalf("expected distinct stable IDs: %#v",docs) }

	after,err:=os.ReadFile(path); if err!=nil { t.Fatal(err) }
	if string(before)!=string(after) { t.Fatal("reader modified source database") }
}
