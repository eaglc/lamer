package letter

import (
    "github.com/eaglc/lamer/codec"
)

type ResponseWriter interface {
    Write(m interface{}) error
    Writer() codec.Codec
}