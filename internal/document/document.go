package document

import "time"

// Document is the normalized, read-only representation of a source record.
type Document struct {
	ID          string         `json:"id"`
	Source      string         `json:"source"`
	SourceID    string         `json:"source_id,omitempty"`
	Title       string         `json:"title,omitempty"`
	Body        string         `json:"body,omitempty"`
	ContentType string         `json:"content_type,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	RawReference string        `json:"raw_reference,omitempty"`
}
