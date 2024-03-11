// Square2
package main

// Square2
type Square2 struct {
	size int // width. height
}

// Rectangle
func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}
