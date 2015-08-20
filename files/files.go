package files

import (
	"errors"
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

func (b Blob) Size() int64 {
	return b.b.Get("size").Int64()
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
	offset, size int64
}

func NewFileReader(f File) *FileReader {
	return &FileReader{
		f,
		js.Global.Get("FileReader").New(),
		0,
		f.Size(),
	}
}

func (fr *FileReader) Close() error {
	fr.fr = nil
	return nil
}

func (fr *FileReader) Read(b []byte) (int, error) {
	n, err := fr.ReadAt(b, int64(fr.offset))
	fr.offset += int64(n)
	return n, err
}

func (fr *FileReader) ReadAt(b []byte, off int64) (int, error) {
	if fr.fr == nil {
		return 0, ErrClosed
	}
	if off >= fr.size {
		return 0, io.EOF
	}

	type readResult struct {
		size int
		err  error
	}

	c := make(chan readResult)
	fr.fr.Set("onloadend", func(*js.Object) {
		arr := js.Global.Get("Uint8Array").New(fr.fr.Get("result"))
		buf := arr.Interface().([]byte)
		go func() {
			if len(buf) == 0 {
				c <- readResult{0, io.EOF}
			} else {
				copy(b, buf)
				c <- readResult{len(buf), nil}
			}
		}()
	})
	e := off + int64(len(b))
	if e > fr.size {
		e = fr.size
	}
	blob := fr.f.b.Call("slice", fr.offset, e)
	fr.fr.Call("readAsArrayBuffer", blob)
	r := <-c
	return r.size, r.err
}

func (fr *FileReader) Seek(offset int64, whence int) (int64, error) {
	if fr.fr == nil {
		return 0, ErrClosed
	}
	switch whence {
	case 0:
		fr.offset = offset
	case 1:
		fr.offset += offset
	case 2:
		fr.offset = fr.size + offset
	}
	if fr.offset < 0 {
		fr.offset = 0
	}
	return fr.offset, nil
}

// Errors
var (
	ErrClosed = errors.New("file closed")
)
