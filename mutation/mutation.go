// Package mutation provides a wrapper for the MutationObserver API
package mutation

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

// Observer wraps a MutationObserver javascript object
type Observer struct {
	*js.Object
}

// New creates a new wrapper around MutationObserver, which is created with the
// given callback
func New(f func([]*Record, *Observer)) *Observer {
	return &Observer{js.Global.Get("MutationObserver").New(f)}
}

// Observe registers a DOM Node to receive events for
func (m *Observer) Observe(n dom.Node, i ObserverInit) {
	j := &observerInit{Object: js.Global.Get("Object").Get("constructor").New()}
	j.ChildList = i.ChildList
	j.Attributes = i.Attributes
	j.CharacterData = i.CharacterData
	j.Subtree = i.Subtree
	j.AttributeOldValue = i.AttributeOldValue
	j.CharacterDataOldValue = i.CharacterDataOldValue
	if i.Attributes {
		j.AttributeFilter = i.AttributeFilter
	}
	m.Call("observe", n.Underlying(), j)
}

// Disconnect stops the observer from receiving events
func (m *Observer) Disconnect() {
	m.Call("disconnect")
}

// TakeRecords empties the record queue, returning what was in it
func (m *Observer) TakeRecords() []*Record {
	o := m.Call("takeRecords")
	l := o.Length()
	mrs := make([]*Record, l)
	for i := 0; i < l; i++ {
		mrs[i] = &Record{Object: o.Index(i)}
	}
	return mrs
}

// Record is a wrapper around a MutationRecord js type
type Record struct {
	*js.Object
	Type               string `js:"type"`
	AttributeName      string `js:"attributeName"`
	AttributeNamespace string `js:"attributeNamespace"`
	OldValue           string `js:"oldValue"`
}

// Target returns the node associated with the record
func (r *Record) Target() dom.Node {
	return dom.WrapNode(r.Get("target"))
}

func wrapNodes(o *js.Object) []dom.Node {
	l := o.Length()
	toRet := make([]dom.Node, l)
	for i := 0; i < l; i++ {
		toRet[i] = dom.WrapNode(o.Index(i))
	}
	return toRet
}

// AddedNodes returns any nodes that were added
func (r *Record) AddedNodes() []dom.Node {
	return wrapNodes(r.Get("addedNodes"))
}

// RemovedNodes returns any nodes that were added
func (r *Record) RemovedNodes() []dom.Node {
	return wrapNodes(r.Get("removedNodes"))
}

// PreviousSibling returns the previous sibling to any added or removed nodes
func (r *Record) PreviousSibling() dom.Node {
	return dom.WrapNode(r.Get("previousSibling"))
}

// NextSibling returns the next sibling to any added or removed nodes
func (r *Record) NextSibling() dom.Node {
	return dom.WrapNode(r.Get("nextSibling"))
}

// ObserverInit contains the options for observing a Node
type ObserverInit struct {
	ChildList             bool
	Attributes            bool
	CharacterData         bool
	Subtree               bool
	AttributeOldValue     bool
	CharacterDataOldValue bool
	AttributeFilter       []string
}

type observerInit struct {
	*js.Object
	ChildList             bool     `js:"childList"`
	Attributes            bool     `js:"attributes"`
	CharacterData         bool     `js:"characterData"`
	Subtree               bool     `js:"subtree"`
	AttributeOldValue     bool     `js:"attributeOldValue"`
	CharacterDataOldValue bool     `js:"characterDataOldValue"`
	AttributeFilter       []string `js:"attributeFilter"`
}
