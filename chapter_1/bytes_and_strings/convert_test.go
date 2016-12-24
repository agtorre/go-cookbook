package bytestrings

import (
	"bytes"
	"testing"
)

func TestToString(t *testing.T) {
	var toStringTests = []struct {
		input   *bytes.Buffer
		output1 string
		output2 error
	}{
		{bytes.NewBufferString("abc"), "abc", nil},
		//{nil, "", nil},
	}

	for _, tt := range toStringTests {
		got1, got2 := toString(tt.input)
		if got1 != tt.output1 || got2 != tt.output2 {
			t.Errorf("got: %#v %#v; want: %#v %#v", got1, got2, tt.output1, tt.output2)
		}
	}
}
