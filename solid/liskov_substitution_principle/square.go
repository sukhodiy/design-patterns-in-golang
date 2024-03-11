// Square
package main

// Square
type Square struct {
	Rectangle
}

// NewSquare
func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// SetWidth
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

// SetHeight
func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}
