package codec

import (
	"io"
)

type NewCodec func(closer io.ReadWriteCloser) Codec

type Codec interface {
	Read(interface{}, ...Option) error
	Write(interface{}, ...Option) error
	Close() error
	String() string
}

type Encoder interface {
	Encode(interface{}, ...Option) error
}

type Decoder interface {
	Decode(interface{}, ...Option) error
}
