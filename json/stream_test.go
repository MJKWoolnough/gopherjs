package json

import (
	"strings"
	"testing"
)

func TestJSONRead(t *testing.T) {
	d := NewDecoder(strings.NewReader("\"Hello\" {} [{\"Pi\": 3.12159E1, \"Meaning Of Life\": 42}, {\"Good?\": true, \"Bad\": false, \"Ugly?\"	:null}] "))
	lengths := []int{7, 3, 89}
	for n, l := range lengths {
		d.readValue()
		if len(d.p.read) != l {
			t.Errorf("test %d: expecting read length of %d, got %d", n+1, l, len(d.p.read))
		}
		d.p.read = d.p.read[:0]
	}
}
