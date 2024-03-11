// Rectangle
package main

// Rectangle
type Rectangle struct {
	width, height int
}

// GetWidth
func (r *Rectangle) GetWidth() int {
	return r.width
}

// SetWidth
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

// GetHeight
func (r *Rectangle) GetHeight() int {
	return r.height
}

// SetHeight
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}
