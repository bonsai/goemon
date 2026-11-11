package cli

import (
	"bytes"
	"strings"
	"testing"

	"goemon/internal/document"
)

func TestBrowseFilterAndDetail(t *testing.T) {
	docs:=[]document.Document{{ID:"1",Source:"jsonl",Title:"Issue Alpha",Body:"reader implementation"}}
	var out bytes.Buffer
	err:=Browse(docs,strings.NewReader("reader\n1\nq\n"),&out)
	if err!=nil { t.Fatal(err) }
	if !strings.Contains(out.String(),"Issue Alpha") || !strings.Contains(out.String(),"reader implementation") { t.Fatalf("missing result/detail: %s",out.String()) }
}
