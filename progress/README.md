# progress
--
    import "github.com/MJKWoolnough/gopherjs/progress"

Package progress creates a simple progress bar in javascript

## Usage

#### type Bar

```go
type Bar struct {
	*dom.HTMLCanvasElement
}
```

Bar is a wrapper around a Canvas element to draw a simple progress bar.

It also implements dom.Node, so it can be Appended to any other node.

#### func  New

```go
func New(fore, back color.Color, width, height int) *Bar
```
New returns a new Bar

#### func (*Bar) Percent

```go
func (b *Bar) Percent(i int)
```
Percent sets the current percentage shown

#### func (*Bar) Reader

```go
func (b *Bar) Reader(r io.Reader, size int) *Reader
```
Reader returns a ProgressReader

#### type Reader

```go
type Reader struct {
	*Bar
	io.Reader
}
```

ProgressReader wraps a ProgressBar to automatically update when an io.Reader is
read

#### func (*Reader) Len

```go
func (r *Reader) Len() int
```
Len returns the total length of the data

#### func (*Reader) Read

```go
func (r *Reader) Read(p []byte) (int, error)
```
Read implements io.Reader
