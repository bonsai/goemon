package eml

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
	"time"

	"goemon/internal/document"
)

type Reader struct { Path string }

func (r Reader) ReadAll(ctx context.Context) ([]document.Document, error) {
	select { case <-ctx.Done(): return nil, ctx.Err(); default: }
	f, err := os.Open(r.Path)
	if err != nil { return nil, err }
	defer f.Close()
	raw, err := io.ReadAll(f)
	if err != nil { return nil, err }
	msg, err := mail.ReadMessage(bytes.NewReader(raw))
	if err != nil { return nil, err }
	body, err := io.ReadAll(msg.Body)
	if err != nil { return nil, err }
	sum := sha256.Sum256(raw)
	var created *time.Time
	if date, e := msg.Header.Date(); e == nil { created = &date }
	meta := map[string]any{
		"from": msg.Header.Get("From"),
		"to": msg.Header.Values("To"),
		"cc": msg.Header.Values("Cc"),
		"message_id": msg.Header.Get("Message-ID"),
	}
	return []document.Document{{
		ID: "eml:" + hex.EncodeToString(sum[:]), Source: "eml",
		SourceID: strings.TrimSuffix(filepath.Base(r.Path), filepath.Ext(r.Path)),
		Title: msg.Header.Get("Subject"), Body: string(body),
		ContentType: msg.Header.Get("Content-Type"), CreatedAt: created,
		Metadata: meta, RawReference: r.Path,
	}}, nil
}
