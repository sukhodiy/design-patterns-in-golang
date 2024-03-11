package main

import (
	"fmt"
)

// Example: print a list of textual items with different formats.

func main() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

	tp.Reset()
	tp.SetOutputFormat(HTML)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)
}
