package json

import (
	"errors"
	"io"
)

type Encoder struct {
	w   io.Writer
	err error
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(v interface{}) error {
	b, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = e.w.Write(b)
	e.err = err
	return err
}

type Decoder struct {
	p   parser
	err error
}

type byteReader struct {
	io.Reader
}

func (br byteReader) ReadByte() (byte, error) {
	b := make([]byte, 1)
	_, err := br.Read(b)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

type parser struct {
	io.ByteReader
	err  error
	buf  byte
	read []byte
}

func (p *parser) Next() byte {
	var b byte
	if p.buf >= 0 {
		b = p.buf
		p.buf = 0
	} else {
		var err error
		b, err = p.ByteReader.ReadByte()
		if err != nil {
			p.err = err
			return 0
		}
	}
	p.read = append(p.read, b)
	return b
}

func (p *parser) Backup() {
	p.buf = p.read[len(p.read)-1]
	p.read = p.read[:len(p.read)-1]
}

func (p *parser) Accept(s string) bool {
	b := p.Next()
	for _, c := range s {
		if c == b {
			return true
		}
	}
	b.Backup()
	return false
}

func (p *parser) AcceptRun(s string) {
	for {
		b := p.Next()
		for _, c := range s {
			if c != b {
				b.Backup()
				return
			}
		}
	}
}

func (p *parser) Peek() byte {
	b := p.Next()
	p.Backup()
	return b
}

func NewDecoder(r io.Reader) *Decoder {
	b, ok := r.(io.ByteReader)
	if !ok {
		b = byteReader{r}
	}
	return &Decoder{p: parser{b}}
}

func (d *Decoder) Decode(v interface{}) error {
	bytes, err := d.readValue()
	if err != nil {
		return err
	}
	if !d.readValue() {
		if d.p.err != nil {
			return d.p.err
		}
		return errInvalidJSON
	}
	s := string(d.p.buf)
	d.p.buf = d.p.buf[:0]
	return UnmarshalString(s, v)
}

const whitespace = " 	\n\r"

func (d *Decoder) readValue() bool {
	d.p.AcceptRun(whitespace)
	var valid bool
	switch d.p.Peek() {
	case '"':
		valid = d.readString()
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		valid = d.readNumber()
	case '{':
		valid = d.readObject()
	case '[':
		valid = d.readArray()
	case 't':
		valid = d.readTrue()
	case 'f':
		valid = d.readFalse()
	case 'n':
		valid = d.readNull()
	}
	return valid
}

func (d *Decoder) readString() bool {
	d.p.Accept("\"")
	for {
		switch c := d.p.Next(); c {
		case '"':
			return true
		case '\\':
			switch d.p.Next() {
			case '"', '\\', '/', 'b', 'n', 'r', 't':
			case 'u':
				if !d.p.Accept("0123456789abcdefABCDEF") || !d.p.Accept("0123456789abcdefABCDEF") || !d.p.Accept("0123456789abcdefABCDEF") || !d.p.Accept("0123456789abcdefABCDEF") {
					return false
				}
			default:
				return false
			}
		case 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x7f:
			return false
		}
	}
}

func (d *Decoder) readNumber(c byte) bool {
	d.p.Accept("-")
	if !d.p.Accept("0") {
		d.p.Accept("123456789")
		d.p.AcceptRun("0123456789")
	}
	if d.p.Accept(".") {
		d.p.AcceptRun("0123456789")
	}
	if d.p.Accept("eE") {
		d.p.Accept("+-")
		d.p.AcceptRun("0123456789")
	}
	return true
}

func (d *Decoder) readObject() bool {

}

func (d *Decoder) readArray() bool {

}

func (d *Decoder) readTrue() bool {
	return d.p.Accept("t") && d.p.Accept("r") && d.p.Accept("u") && d.p.Accept("e")
}

func (d *Decoder) readFalse() bool {
	return d.p.Accept("f") && d.p.Accept("a") && d.p.Accept("l") && d.p.Accept("s") && d.p.Accept("e")
}

func (d *Decoder) readNull() bool {
	return d.p.Accept("n") && d.p.Accept("u") && d.p.Accept("l") && d.p.Accept("l")
}

var errInvalidJSON = errors.New("invalid JSON")
