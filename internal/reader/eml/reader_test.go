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


func TestReadAllPreservesRepeatedRecipientHeaders(t *testing.T) {
	path := filepath.Join(t.TempDir(), "repeated.eml")
	raw := "From: sender@example.com\r\n" +
		"To: first@example.com\r\n" +
		"To: second@example.com\r\n" +
		"Cc: copy-one@example.com\r\n" +
		"Cc: copy-two@example.com\r\n" +
		"Subject: Repeated recipients\r\n\r\nBody"
	if err := os.WriteFile(path, []byte(raw), 0600); err != nil {
		t.Fatal(err)
	}

	docs, err := (Reader{Path: path}).ReadAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(docs) != 1 {
		t.Fatalf("got %d documents, want 1", len(docs))
	}

	to, ok := docs[0].Metadata["to"].([]string)
	if !ok {
		t.Fatalf("metadata to has type %T, want []string", docs[0].Metadata["to"])
	}
	if len(to) != 2 || to[0] != "first@example.com" || to[1] != "second@example.com" {
		t.Fatalf("unexpected To values: %#v", to)
	}

	cc, ok := docs[0].Metadata["cc"].([]string)
	if !ok {
		t.Fatalf("metadata cc has type %T, want []string", docs[0].Metadata["cc"])
	}
	if len(cc) != 2 || cc[0] != "copy-one@example.com" || cc[1] != "copy-two@example.com" {
		t.Fatalf("unexpected Cc values: %#v", cc)
	}
}

func TestReadAllAllowsMissingRecipientHeaders(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing-recipients.eml")
	raw := "From: sender@example.com\r\nSubject: No recipients\r\n\r\nBody"
	if err := os.WriteFile(path, []byte(raw), 0600); err != nil {
		t.Fatal(err)
	}

	docs, err := (Reader{Path: path}).ReadAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(docs) != 1 {
		t.Fatalf("got %d documents, want 1", len(docs))
	}
	for _, key := range []string{"to", "cc"} {
		values, ok := docs[0].Metadata[key].([]string)
		if !ok {
			t.Fatalf("metadata %s has type %T, want []string", key, docs[0].Metadata[key])
		}
		if len(values) != 0 {
			t.Errorf("metadata %s = %#v, want no values", key, values)
		}
	}
}
