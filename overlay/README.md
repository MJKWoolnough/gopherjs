# overlay
--
    import "github.com/MJKWoolnough/gopherjs/overlay"

Package overlay creates a simple 'window' overlay system for gopherjs

## Usage

#### type Overlay

```go
type Overlay struct {
	dom.Element
}
```

Overlay is a dom.Element which

#### func  New

```go
func New(e dom.Node) *Overlay
```
New wraps a dom.Element in an Overlay

#### func (*Overlay) Close

```go
func (o *Overlay) Close() error
```
Close closes the overlay, and runs any given onClose handle funcs

#### func (*Overlay) OnClose

```go
func (o *Overlay) OnClose(f func())
```
OnClose adds an onClose handler for the overlay
