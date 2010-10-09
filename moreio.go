package moreio

import (
	"io"
	"os"
	)

type multiplex struct {
	r io.ReadCloser
	w io.WriteCloser
}

func NewReadWriteCloser(r io.ReadCloser, w io.WriteCloser) *multiplex {
	return &multiplex{r: r, w: w}
}

func (m *multiplex) Read(p []byte) (n int, err os.Error) {
	return m.r.Read(p)
}

func (m *multiplex) Write(p []byte) (n int, err os.Error) {
	return m.w.Write(p)
}

func (m *multiplex) Close() os.Error {
	err := m.w.Close()
	err2 := m.r.Close()
	if err == nil {
		err = err2
	}
	return err
}
