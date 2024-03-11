// Journal demonstrates wrong approach that breaks SRP, because Journal
// struct has too many responsibilities.
package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// Entries counter
var entryCount = 0

// Journal
type Journal struct {
	entries []string
}

// Return journal entries as a string
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Add entry to journal
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// Remove entry from journal
func (j *Journal) RemoveEntry(index int) {
	// ...
}

// Separation of concerns below:

// Save journal content to file
func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

// Load journal content from file
func (j *Journal) Load(filename string) {
	//
}

// Load journal content from web
func (j *Journal) LoadFromWeb(url *url.URL) {
	//
}
