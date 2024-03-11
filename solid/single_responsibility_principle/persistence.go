// Persistence
package main

import (
	"os"
	"strings"
)

// Persistence
type Persistence struct {
	lineSeparator string
}

// Save journal content to file
func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename,
		[]byte(strings.Join(j.entries, LineSeparator)),
		0644)
}
