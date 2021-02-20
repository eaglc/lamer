package codec

import "io"

type Options struct {
	r io.Reader
	w io.Writer
	c io.Closer
	o io.ReadWriteCloser
}

type Option func(*Options)

func Reader(r io.Reader) Option {
	return func(o *Options) {
		o.r = r
	}
}

func Writer(w io.Writer) Option {
	return func(o *Options) {
		o.w = w
	}
}

func Closer(c io.Closer) Option {
	return func(o *Options) {
		o.c = c
	}
}

func ReadWriteCloser(io io.ReadWriteCloser) Option {
	return func(o *Options) {
		o.o = io
	}
}
