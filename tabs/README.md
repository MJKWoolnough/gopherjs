# tabs
--
    import "github.com/MJKWoolnough/gopherjs/tabs"

Package tabs implements a simple tab interface for HTML documents

## Usage

#### func  New

```go
func New(t []Tab) dom.Node
```
New takes a list of tabs and generates a tabbed interface, which is return as a
document fragment

#### type Tab

```go
type Tab struct {
	// The name of the Tab
	Name string
	// Func to run when the tab is selected
	Func func(dom.Element)
}
```

Tab represents a tab in a menu
