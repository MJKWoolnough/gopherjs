# mutation
--
    import "github.com/MJKWoolnough/gopherjs/mutation"

Package mutation provides a wrapper for the MutationObserver API

## Usage

#### type Observer

```go
type Observer struct {
	*js.Object
}
```

Observer wraps a MutationObserver javascript object

#### func  New

```go
func New(f func([]*Record, *Observer)) *Observer
```
New creates a new wrapper around MutationObserver, which is created with the
given callback

#### func (*Observer) Disconnect

```go
func (m *Observer) Disconnect()
```
Disconnect stops the observer from receiving events

#### func (*Observer) Observe

```go
func (m *Observer) Observe(n dom.Node, i ObserverInit)
```
Observe registers a DOM Node to receive events for

#### func (*Observer) TakeRecords

```go
func (m *Observer) TakeRecords() []*Record
```
TakeRecords empties the record queue, returning what was in it

#### type ObserverInit

```go
type ObserverInit struct {
	ChildList             bool
	Attributes            bool
	CharacterData         bool
	Subtree               bool
	AttributeOldValue     bool
	CharacterDataOldValue bool
	AttributeFilter       []string
}
```

ObserverInit contains the options for observing a Node

#### type Record

```go
type Record struct {
	*js.Object
	Type               string `js:"type"`
	AttributeName      string `js:"attributeName"`
	AttributeNamespace string `js:"attributeNamespace"`
	OldValue           string `js:"oldValue"`
}
```

Record is a wrapper around a MutationRecord js type

#### func (*Record) AddedNodes

```go
func (r *Record) AddedNodes() []dom.Node
```
AddedNodes returns any nodes that were added

#### func (*Record) NextSibling

```go
func (r *Record) NextSibling() dom.Node
```
NextSibling returns the next sibling to any added or removed nodes

#### func (*Record) PreviousSibling

```go
func (r *Record) PreviousSibling() dom.Node
```
PreviousSibling returns the previous sibling to any added or removed nodes

#### func (*Record) RemovedNodes

```go
func (r *Record) RemovedNodes() []dom.Node
```
RemovedNodes returns any nodes that were added

#### func (*Record) Target

```go
func (r *Record) Target() dom.Node
```
Target returns the node associated with the record
