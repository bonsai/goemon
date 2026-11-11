package reader

import (
	"context"
	"goemon/internal/document"
)

// Reader exposes source records as normalized documents. Implementations must
// not mutate the source.
type Reader interface {
	ReadAll(context.Context) ([]document.Document, error)
}
