package moreio

import "io"

type multiplex struct {
	io.ReadCloser
	io.WriteCloser
}

func NewReadWriteCloser(r io.ReadCloser, w io.WriteCloser) *multiplex {
	return &multiplex{ReadCloser: r, WriteCloser: w}
}

func (m *multiplex) Close() error {
	err := m.WriteCloser.Close()
	err2 := m.ReadCloser.Close()
	if err == nil {
		err = err2
	}
	return err
}
