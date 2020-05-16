package logging

import (
	"bytes"
	"io"
)

type writer struct {
	bytes.Buffer
	out io.Writer
}

func (w *writer) Flush() error {
	b := w.Buffer.Bytes()
	_, err := w.out.Write(b)
	w.Buffer.Reset()
	return err
}
