package session

import "github.com/eaglc/lamer/codec"

type Session interface {
    Codec() codec.Codec

    // metadata for codec
    Metadata() interface{}
}