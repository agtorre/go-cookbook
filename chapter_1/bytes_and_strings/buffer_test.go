package bytestrings

import (
	"bytes"
	"testing"
)

func TestBuffer(t *testing.T) {
	var bufferTests = []struct {
		input  string
		output *bytes.Buffer
	}{
		{"abc", bytes.NewBufferString("abc")},
		{"", bytes.NewBufferString("")},
		{"❤️", bytes.NewBufferString("❤️")},
	}

	for _, tt := range bufferTests {
		if got, want := Buffer(tt.input), tt.output; got.String() != want.String() {
			t.Errorf("got: %#v; want: %#v", got, want)
		}
	}
}
