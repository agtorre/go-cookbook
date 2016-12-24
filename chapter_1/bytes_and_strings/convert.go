package bytestrings

import (
	"io"
	"io/ioutil"
)

// toString is an example of taking an io.Reader and consuming it all,
// then returning a string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
