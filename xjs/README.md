# xjs
--
    import "github.com/MJKWoolnough/gopherjs/xjs"

Package xjs provides some simple, but often needed shortcut funcs for gopherJS

## Usage

#### func  Alert

```go
func Alert(format string, params ...interface{})
```
Alert provides for formated alert boxes

#### func  AppendChildren

```go
func AppendChildren(parent dom.Node, children ...dom.Node) dom.Node
```
AppendChildren appends all the given children to the parent.

#### func  Body

```go
func Body() *dom.HTMLBodyElement
```
Get the document's body element

#### func  CreateElement

```go
func CreateElement(name string) dom.Element
```
CreateElement is a shortcut to create an element with the given name

#### func  DocumentFragment

```go
func DocumentFragment() *dom.BasicHTMLElement
```
DocumentFragment returns a new DocumentFragment as a dom.Node

#### func  Log

```go
func Log(format string, params ...interface{})
```

#### func  RemoveChildren

```go
func RemoveChildren(node dom.Node) dom.Node
```
RemoveChildren removes all of the child nodes of the node given

#### func  SetInnerText

```go
func SetInnerText(node dom.Node, text string) dom.Node
```
SetInnerText removes all child nodes from the given node and sets a single Text
Node with the given string

#### func  SetPreText

```go
func SetPreText(node dom.Node, text string) dom.Node
```
SetPreText does similar to SetInnerText, but linebreaks are converted to <br />s

#### func  Text

```go
func Text(text string) *dom.Text
```
TextNode creates a text node containing the givin text
