package eml

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestReadAllParsesBasicMessage(t *testing.T) {
	path:=filepath.Join(t.TempDir(),"message.eml")
	raw:="From: sender@example.com\r\nTo: reader@example.com\r\nSubject: Hello\r\nDate: Tue, 10 Nov 2026 12:00:00 +0000\r\nMessage-ID: <abc@example.com>\r\nContent-Type: text/plain; charset=utf-8\r\n\r\nMessage body"
	if err:=os.WriteFile(path,[]byte(raw),0600); err!=nil { t.Fatal(err) }
	docs,err:=(Reader{Path:path}).ReadAll(context.Background())
	if err!=nil { t.Fatal(err) }
	if len(docs)!=1 || docs[0].Title!="Hello" || docs[0].Body!="Message body" { t.Fatalf("unexpected result: %#v",docs) }
	if docs[0].Metadata["from"]!="sender@example.com" { t.Fatalf("missing sender: %#v",docs[0].Metadata) }
}
