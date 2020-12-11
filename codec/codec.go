package codec

import (
    "io"
)

type NewCodec func(closer io.ReadWriteCloser) Codec

type Codec interface {
    Encoder
    Decoder
    Close() error
    String() string
}

type Encoder interface {
    Encode(interface{}) error
}

type Decoder interface {
    Decode() (interface{}, error)
}