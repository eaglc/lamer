package session

import "github.com/eaglc/lamer/codec"

type State int32

type Session interface {
    // codec session use
    Codec() codec.Codec

    // metadata
    Metadata() interface{}

    // state of this session
    State() State

    // error of this session
    Error() error
}