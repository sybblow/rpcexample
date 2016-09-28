package rpcexample

import "io"

type LineWriter struct {
	m io.WriteCloser
}

func NewLineWriter(wr io.WriteCloser) *LineWriter {
	return &LineWriter{wr}
}

func (wr *LineWriter) Write(p []byte) (n int, err error) {
	n, err = wr.m.Write(p)
	if err != nil {
		return
	}
	nn, err := io.WriteString(wr.m, "\n")
	if err != nil {
		return
	}
	n += nn
	return
}

func (wr *LineWriter) Close() error {
	return wr.m.Close()
}
