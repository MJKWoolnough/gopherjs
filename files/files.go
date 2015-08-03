package main

import (
	"io"
	"time"

	"honnef.co/go/js/dom"

	"github.com/gopherjs/gopherjs/js"
)

type Blob struct {
	b *js.Object
}

func NewBlob(b *js.Object) Blob {
	return Blob{b}
}

func (b Blob) Size() int {
	return b.b.Get("size").Int()
}

func (b Blob) Type() string {
	return b.b.Get("type").String()
}

func (b Blob) Slice(i, j int) Blob {
	return NewBlob(b.b.Call("slice", i, j))
}

type File struct {
	Blob
}

func NewFile(f *dom.File) File {
	return File{NewBlob(f.Object)}
}

func (f File) LastModifiedDate() time.Time {
	return f.b.Get("lastModifiedDate").Interface().(time.Time)
}

func (f File) Name() string {
	return f.b.Get("name").String()
}

type FileReader struct {
	f            File
	fr           *js.Object
	offset, size int
}

func NewFileReader(f File) *FileReader {
	return &FileReader{
		f,
		js.Global.Get("FileReader").New(),
		0,
		f.Size(),
	}
}

func (fr *FileReader) Read(p []byte) (int, error) {
	if fr.offset == fr.size {
		return 0, io.EOF
	}
	type result struct {
		size int
		err  error
	}
	c := make(chan result)
	fr.fr.Set("onloadend", func(*js.Object) {
		arr := js.Global.Get("Uint8Array").New(fr.fr.Get("result"))
		buf := arr.Interface().([]byte)
		go func() {
			if len(buf) == 0 {
				c <- result{0, io.EOF}
			} else {
				copy(p, buf)
				c <- result{len(buf), nil}
			}
		}()
	})
	e := fr.offset + len(p)
	if e > fr.size {
		e = fr.size
	}
	blob := fr.f.b.Call("slice", fr.offset, e)
	fr.fr.Call("readAsArrayBuffer", blob)
	res := <-c
	fr.offset = e
	return res.size, res.err
}
