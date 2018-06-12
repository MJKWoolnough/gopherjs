// Package progress creates a simple progress bar in javascript
package progress // import "vimagination.zapto.org/gopherjs/progress"

import (
	"image/color"
	"io"
	"strconv"

	"vimagination.zapto.org/gopherjs/xdom"

	"honnef.co/go/js/dom"
)

// Bar is a wrapper around a Canvas element to draw a simple progress
// bar.
//
// It also implements dom.Node, so it can be Appended to any other node.
type Bar struct {
	*dom.HTMLCanvasElement
	percent                int
	ctx                    *dom.CanvasRenderingContext2D
	width, height          int
	foreground, background string
}

// Percent sets the current percentage shown
func (b *Bar) Percent(i int) {
	if i != b.percent {
		b.percent = i
		b.draw()
	}
}

func (b *Bar) draw() {
	b.ctx.FillStyle = b.background
	b.ctx.FillRect(0, 0, b.width, b.height)
	b.ctx.FillStyle = b.foreground
	b.ctx.FillRect(0, 0, b.width*b.percent/100, b.height)
}

// New returns a new Bar
func New(fore, back color.Color, width, height int) *Bar {
	c := xdom.Canvas()
	c.Width = width
	c.Height = height
	r, g, b, _ := fore.RGBA()
	foreground := "rgb(" + strconv.Itoa(int(r>>8)) + ", " + strconv.Itoa(int(g>>8)) + ", " + strconv.Itoa(int(b>>8)) + ")"
	r, g, b, _ = back.RGBA()
	background := "rgb(" + strconv.Itoa(int(r>>8)) + ", " + strconv.Itoa(int(g>>8)) + ", " + strconv.Itoa(int(b>>8)) + ")"
	bar := &Bar{
		HTMLCanvasElement: c,
		ctx:               c.GetContext2d(),
		width:             width,
		height:            height,
		foreground:        foreground,
		background:        background,
	}
	bar.draw()
	return bar
}

// Reader wraps a ProgressBar to automatically update when an io.Reader
// is read
type Reader struct {
	*Bar
	io.Reader
	offset, size int
}

// Reader returns a Reader type
func (b *Bar) Reader(r io.Reader, size int) *Reader {
	return &Reader{b, r, 0, size}
}

// Read implements io.Reader
func (r *Reader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	r.offset += n
	r.Percent(100 * r.offset / r.size)
	return n, err
}

// Len returns the total length of the data
func (r *Reader) Len() int {
	return int(r.size)
}
