# Goemon Reader

Goemon is a read-only, multi-source reader for mail, EML, SQLite, JSONL and web documents. It normalizes source records into a common Document shape so CLI and future UI/API clients can use the same data model.

Agent orchestration, worker lifecycle management and image generation are intentionally outside this repository's target responsibility. Those capabilities belong in `bons.ai-swarm`.

## Status

The Reader foundation currently includes:
- Common `document.Document` type and `reader.Reader` contract
- EML file reader
- JSONL streaming reader
- SQLite read-only reader
- CLI output as JSONL

Web reading, interactive fuzzy selection, full-text indexing, richer MIME attachment extraction and a dedicated HTTP API remain follow-up work.

## Commands

```sh
go run ./src/cmd/goemon read eml ./mail/example.eml
go run ./src/cmd/goemon read jsonl ./data/documents.jsonl
go run ./src/cmd/goemon read sqlite ./data/archive.db
go run ./src/cmd/goemon read sqlite ./data/archive.db --table messages --limit 200
go run ./src/cmd/goemon export jsonl --input ./data/documents.jsonl
```

All commands emit one normalized JSON document per output line. Errors are written to stderr and return a non-zero exit code.

## Document shape

| Field | Meaning |
|---|---|
| `id` | Stable or content-derived identifier |
| `source` | Source adapter name |
| `source_id` | Identifier within the source |
| `title` | Human-readable title |
| `body` | Text or serialized record body |
| `content_type` | MIME/content type |
| `created_at`, `updated_at` | Source timestamps, when available |
| `metadata` | Source-specific fields |
| `raw_reference` | Path or source locator; original data is not modified |

## Design constraints

- Readers never mutate source data.
- SQLite is opened in read-only/query-only mode.
- JSONL is processed line by line; malformed input reports its line number.
- Secrets must come from external configuration, never source constants.
- Raw source records remain the authority; normalized Documents are a read model.

## Development

```sh
go test ./...
go build ./src/cmd/goemon
```

See [docs/architecture.md](docs/architecture.md) for the intended package boundaries and [docs/adr/](docs/adr/) for decisions.
