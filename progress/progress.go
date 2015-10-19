// Package progress creates a simple progress bar in javascript
package progress

import (
	"image/color"
	"io"
	"strconv"

	"github.com/MJKWoolnough/gopherjs/xdom"

	"honnef.co/go/js/dom"
)

// ProgressBar is a wrapper around a Canvas element to draw a simple progress
// bar.
//
// It also implements dom.Node, so it can be Appended to any other node.
type ProgressBar struct {
	*dom.HTMLCanvasElement
	percent                int
	ctx                    *dom.CanvasRenderingContext2D
	width, height          int
	foreground, background string
}

// Percent sets the current percentage shown
func (p *Bar) Percent(i int) {
	if i != p.percent {
		p.percent = i
		p.draw()
	}
}

func (p *ProgressBar) draw() {
	p.ctx.FillStyle = p.background
	p.ctx.FillRect(0, 0, p.width, p.height)
	p.ctx.FillStyle = p.foreground
	p.ctx.FillRect(0, 0, p.width*p.percent/100, p.height)
}

// New returns a new ProgressBar
func New(fore, back color.Color, width, height int) *ProgressBar {
	c := xdom.Canvas()
	c.Width = width
	c.Height = height
	r, g, b, _ := fore.RGBA()
	foreground := "rgb(" + strconv.Itoa(int(r>>8)) + ", " + strconv.Itoa(int(g>>8)) + ", " + strconv.Itoa(int(b>>8)) + ")"
	r, g, b, _ = back.RGBA()
	background := "rgb(" + strconv.Itoa(int(r>>8)) + ", " + strconv.Itoa(int(g>>8)) + ", " + strconv.Itoa(int(b>>8)) + ")"
	pb := &ProgressBar{
		HTMLCanvasElement: c,
		ctx:               c.GetContext2d(),
		width:             width,
		height:            height,
		foreground:        foreground,
		background:        background,
	}
	pb.draw()
	return pb
}

// ProgressReader wraps a ProgressBar to automatically update when an io.Reader
// is read
type ProgressReader struct {
	*ProgressBar
	io.Reader
	offset, size int
}

// Reader returns a ProgressReader
func (p *ProgressBar) Reader(r io.Reader, size int) *ProgressReader {
	return &ProgressReader{p, r, 0, size}
}

// Read implements io.Reader
func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.offset += n
	pr.Percent(100 * pr.offset / pr.size)
	return n, err
}

// Len returns the total length of the data
func (pr *ProgressReader) Len() int {
	return int(pr.size)
}
