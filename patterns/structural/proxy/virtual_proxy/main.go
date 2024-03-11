package main

import (
	"fmt"
)

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{filename: filename}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func main() {
	// we still load the file, but we don't draw it.
	// How not to load a file until we draw it?
	//_ = NewBitmap("demo.png")

	bmp := NewLazyBitmap("demo.png")
	DrawImage(bmp)
	DrawImage(bmp)
}
