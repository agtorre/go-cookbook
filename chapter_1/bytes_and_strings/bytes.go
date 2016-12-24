package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

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

// toString is an example of taking an io.Reader and consuming it all,
// then returning a string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// WorkWithBuffer will make use of the buffer created by the Buffer function
func WorkWithBuffer() error {
	rawString := "it's easy to encode a string or unicode into a byte array ❤️"

	b := Buffer(rawString)

	// we can quickly convert a buffer back into byes with
	// b.Bytes() or a string with b.String()
	fmt.Println(b.String())

	// because this is an io Reader we can make use of generic io reader functions such as
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// we can also take our bytes and create a bytes reader
	// these readers implement io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner interfaces
	reader := bytes.NewReader([]byte(rawString))

	// we can also plug it into a scanner that allows buffered reading and tokenzation
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// iterate over all of the scan events
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return nil
}
