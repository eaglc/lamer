package codec

import (
	"bufio"
	"bytes"
	"io"
)

type Options struct {
	Reader io.Reader
	Writer io.Writer
	Closer io.Closer
	RWC    io.ReadWriteCloser

	Buffer *bytes.Buffer

	ReadBuffer  *bufio.Reader
	WriteBuffer *bufio.Writer
}

type Option func(*Options)

func Reader(r io.Reader) Option {
	return func(o *Options) {
		o.Reader = r
	}
}

func Writer(w io.Writer) Option {
	return func(o *Options) {
		o.Writer = w
	}
}

func Closer(c io.Closer) Option {
	return func(o *Options) {
		o.Closer = c
	}
}

func ReadWriteCloser(io io.ReadWriteCloser) Option {
	return func(o *Options) {
		o.RWC = io
	}
}

func Buffer(buf *bytes.Buffer) Option {
	return func(o *Options) {
		o.Buffer = buf
	}
}

func ReadBuffer(buf *bufio.Reader) Option {
	return func(o *Options) {
		o.ReadBuffer = buf
	}
}

func WriteBuffer(buf *bufio.Writer) Option {
	return func(o *Options) {
		o.WriteBuffer = buf
	}
}
