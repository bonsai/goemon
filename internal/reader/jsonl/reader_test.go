package jsonl

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadAllAssignsIDAndPreservesSource(t *testing.T) {
	path := filepath.Join(t.TempDir(), "input.jsonl")
	if err := os.WriteFile(path, []byte("{\"title\":\"hello\",\"metadata\":{\"tag\":\"x\"}}\n"), 0600); err != nil { t.Fatal(err) }
	docs, err := (Reader{Path:path}).ReadAll(context.Background())
	if err != nil { t.Fatal(err) }
	if len(docs)!=1 || docs[0].ID=="" || docs[0].Source!="jsonl" || docs[0].Title!="hello" { t.Fatalf("unexpected documents: %#v",docs) }
	if docs[0].Metadata["tag"]!="x" { t.Fatalf("metadata lost: %#v",docs[0].Metadata) }
}

func TestReadAllReportsBadLine(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.jsonl")
	if err := os.WriteFile(path, []byte("{}\nnot-json\n"), 0600); err != nil { t.Fatal(err) }
	_, err := (Reader{Path:path}).ReadAll(context.Background())
	if err==nil || !strings.Contains(err.Error(), ":2:") { t.Fatalf("expected line 2 error, got %v",err) }
}
