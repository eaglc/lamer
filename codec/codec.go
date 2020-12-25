package codec

import (
    "io"
)

type NewCodec func(closer io.ReadWriteCloser) Codec

type Codec interface {
    Read(interface{}) error
    Write(interface{}) error
    Close() error
    String() string
}

type Encoder interface {
    Encode(interface{}) error
}

type Decoder interface {
    Decode(interface{}) error
}