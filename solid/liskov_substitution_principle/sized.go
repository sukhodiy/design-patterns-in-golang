// Sized
package main

// Main interface to support by figures
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}
