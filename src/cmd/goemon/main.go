package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"goemon/internal/document"
	"goemon/internal/reader"
	"goemon/internal/reader/eml"
	jsonlreader "goemon/internal/reader/jsonl"
	sqlitereader "goemon/internal/reader/sqlite"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	ctx := context.Background()
	var docs []document.Document
	var err error

	switch os.Args[1] {
	case "read":
		if len(os.Args) < 4 {
			usage()
			os.Exit(2)
		}
		switch os.Args[2] {
		case "eml":
			docs, err = (eml.Reader{Path: os.Args[3]}).ReadAll(ctx)
		case "jsonl":
			docs, err = (jsonlreader.Reader{Path: os.Args[3]}).ReadAll(ctx)
		case "sqlite":
			fs := flag.NewFlagSet("sqlite", flag.ExitOnError)
			table := fs.String("table", "", "table name")
			limit := fs.Int("limit", 1000, "maximum rows per table")
			_ = fs.Parse(os.Args[4:])
			docs, err = (sqlitereader.Reader{Path: os.Args[3], Table: *table, Limit: *limit}).ReadAll(ctx)
		default:
			usage()
			os.Exit(2)
		}
	case "export":
		if len(os.Args) < 3 || os.Args[2] != "jsonl" {
			usage()
			os.Exit(2)
		}
		fs := flag.NewFlagSet("export", flag.ExitOnError)
		input := fs.String("input", "", "input JSONL")
		_ = fs.Parse(os.Args[3:])
		if *input == "" {
			fmt.Fprintln(os.Stderr, "export jsonl requires --input")
			os.Exit(2)
		}
		docs, err = (jsonlreader.Reader{Path: *input}).ReadAll(ctx)
	default:
		usage()
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "goemon:", err)
		os.Exit(1)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	for _, d := range docs {
		if err := enc.Encode(d); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "goemon read eml <file> | read jsonl <file> | read sqlite <db> [--table name] [--limit n] | export jsonl --input <file>")
}

var _ reader.Reader = eml.Reader{}
var _ reader.Reader = jsonlreader.Reader{}
var _ reader.Reader = sqlitereader.Reader{}
