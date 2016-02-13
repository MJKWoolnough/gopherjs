package json

import "io"

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
	r   io.Reader
	err error
}

func NewDecoder(r io.Reader) *Decoder {
	return Decoder{r: r}
}

func (d *Decoder) Decode(v interface{}) error {
	// read single json object from reader. Until either io.EOF or an invalid combination of chars, e.g. }{, ][, ]{, }[, etc.
	str := ""
	return Unmarshal(str, v)
}
