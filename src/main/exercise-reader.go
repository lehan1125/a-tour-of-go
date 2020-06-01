package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	blength := len(b)
	for i := range b {
		b[i] = 'A'
	}
	return blength, nil
}
func main() {
	reader.Validate(MyReader{})
}
