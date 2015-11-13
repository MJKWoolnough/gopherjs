package mutation

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type Observer struct {
	*js.Object
}

func New(f func([]*Record, *Observer)) *Observer {
	return &Observer{js.Global.Get("MutationObserver").New(f)}
}

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

func (m *Observer) Disconnect() {
	m.Call("disconnect")
}

func (m *Observer) TakeRecords() []*Record {
	o := m.Call("takeRecords")
	l := o.Length()
	mrs := make([]*Record, l)
	for i := 0; i < l; i++ {
		mrs[i] = &Record{Object: o.Index(i)}
	}
	return mrs
}

type Record struct {
	*js.Object
	Type               string `js:"type"`
	AttributeName      string `js:"attributeName"`
	AttributeNamespace string `js:"attributeNamespace"`
	OldValue           string `js:"oldValue"`
}

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

func (r *Record) AddedNodes() []dom.Node {
	return wrapNodes(r.Get("addedNodes"))
}

func (r *Record) RemovedNodes() []dom.Node {
	return wrapNodes(r.Get("removedNodes"))
}

func (r *Record) PreviousSibling() dom.Node {
	return dom.WrapNode(r.Get("previousSibling"))
}

func (r *Record) NextSibling() dom.Node {
	return dom.WrapNode(r.Get("nextSibling"))
}

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
