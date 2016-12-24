package bytestrings

import "bytes"

// Buffer demonstrates some tricks for initializing bytes Buffers
// These buffers implement an io.Reader interface
func Buffer(rawString string) *bytes.Buffer {

	// we'll start with a string encoded into raw bytes
	rawBytes := []byte(rawString)

	// there are a number of ways to create a buffer from the raw bytes or from the original string
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// alternatively
	b = bytes.NewBuffer(rawBytes)

	// and avoiding the intial byte array altogether
	b = bytes.NewBufferString(rawString)

	return b
}
