package main

import "strings"

type HTMLListStrategy struct{}

func (h *HTMLListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HTMLListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h *HTMLListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("  <li>" + item + "</li>\n")
}
