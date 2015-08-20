package progress

import (
	"image/color"
	"io"
	"strconv"

	"github.com/MJKWoolnough/gopherjs/xjs"

	"honnef.co/go/js/dom"
)

type ProgressBar struct {
	*dom.HTMLCanvasElement
	percent                int
	ctx                    *dom.CanvasRenderingContext2D
	width, height          int
	foreground, background string
}

func (p *ProgressBar) Percent(i int) {
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

func New(fore, back color.Color, width, height int) *ProgressBar {
	c := xjs.CreateElement("canvas").(*dom.HTMLCanvasElement)
	c.SetAttribute("width", strconv.Itoa(width))
	c.SetAttribute("height", strconv.Itoa(height))
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

type ProgressReader struct {
	*ProgressBar
	io.Reader
	offset, size int64
}

func (p *ProgressBar) Reader(r io.Reader, size int) *ProgressReader {
	return &ProgressReader{p, r, 0, int64(size)}
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.offset += int64(n)
	pr.Percent(int(100 * pr.offset / pr.size))
	return n, err
}

func (pr *ProgressReader) Size() int64 {
	return pr.size
}
