package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"goemon/internal/document"
)

// Browse presents a keyboard-friendly filter/list/detail loop. It deliberately
// uses line input so it works in terminals, pipes and Windows consoles without
// requiring a platform-specific raw-terminal dependency.
func Browse(docs []document.Document, in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, "filter (blank = all, q = quit): ")
		if !scanner.Scan() { return scanner.Err() }
		query := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(query, "q") { return nil }
		matches := filter(docs, query)
		sort.SliceStable(matches, func(i,j int) bool { return matches[i].Title < matches[j].Title })
		if len(matches)==0 { fmt.Fprintln(out, "no matches"); continue }
		for i,d := range matches {
			title:=d.Title; if title=="" { title="(untitled)" }
			fmt.Fprintf(out, "%3d  %-12s  %s\n", i+1, d.Source, title)
		}
		fmt.Fprint(out, "select number (Enter = search again, q = quit): ")
		if !scanner.Scan() { return scanner.Err() }
		choice:=strings.TrimSpace(scanner.Text())
		if strings.EqualFold(choice,"q") { return nil }
		if choice=="" { continue }
		var index int
		if _,err:=fmt.Sscanf(choice,"%d",&index); err!=nil || index<1 || index>len(matches) {
			fmt.Fprintln(out,"invalid selection")
			continue
		}
		showDocument(out,matches[index-1])
	}
}

func filter(docs []document.Document, query string) []document.Document {
	if query=="" { return append([]document.Document(nil),docs...) }
	q:=strings.ToLower(query)
	out:=make([]document.Document,0)
	for _,d:=range docs {
		haystack:=strings.ToLower(strings.Join([]string{d.Title,d.Body,d.Source,d.SourceID}, "\n"))
		if strings.Contains(haystack,q) { out=append(out,d) }
	}
	return out
}

func showDocument(out io.Writer,d document.Document) {
	fmt.Fprintln(out,"\n--- DOCUMENT ---")
	fmt.Fprintf(out,"ID: %s\nSource: %s\nTitle: %s\nContent-Type: %s\n",d.ID,d.Source,d.Title,d.ContentType)
	fmt.Fprintf(out,"Created: %v\nUpdated: %v\n",d.CreatedAt,d.UpdatedAt)
	fmt.Fprintln(out,"Body:")
	fmt.Fprintln(out,d.Body)
	if len(d.Metadata)>0 { fmt.Fprintln(out,"Metadata:"); for k,v:=range d.Metadata { fmt.Fprintf(out,"  %s: %v\n",k,v) } }
	fmt.Fprintf(out,"Raw: %s\n--------------\n\n",d.RawReference)
}

var _ = os.Stdin
