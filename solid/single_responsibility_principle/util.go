// Util
package main

import (
	"os"
	"strings"
)

// LineSeparator
var LineSeparator = "\n"

// Save journal content to file
func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename,
		[]byte(strings.Join(j.entries, LineSeparator)),
		0644)
}
