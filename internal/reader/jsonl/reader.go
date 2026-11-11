package jsonl

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"goemon/internal/document"
)

type Reader struct { Path string }

func (r Reader) ReadAll(ctx context.Context) ([]document.Document, error) {
	f, err := os.Open(r.Path)
	if err != nil { return nil, err }
	defer f.Close()
	out := make([]document.Document, 0)
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 64*1024), 16*1024*1024)
	line := 0
	for scanner.Scan() {
		line++
		select { case <-ctx.Done(): return nil, ctx.Err(); default: }
		raw := strings.TrimSpace(scanner.Text())
		if raw == "" { continue }
		var d document.Document
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			return nil, fmt.Errorf("%s:%d: %w", r.Path, line, err)
		}
		if d.ID == "" {
			sum := sha256.Sum256([]byte(raw))
			d.ID = "jsonl:" + hex.EncodeToString(sum[:])
		}
		if d.Source == "" { d.Source = "jsonl" }
		if d.RawReference == "" { d.RawReference = fmt.Sprintf("%s:%d", r.Path, line) }
		out = append(out, d)
	}
	if err := scanner.Err(); err != nil { return nil, err }
	return out, nil
}
