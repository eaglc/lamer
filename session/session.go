package session

import (
    "github.com/eaglc/lamer/codec"
)

type State int32

type Session interface {
    // codec session use
    Codec() codec.Codec
    SetCodec(codec.Codec)

    // metadata
    Metadata() interface{}
    SetMetadata(interface{})

    // state of this session
    State() State
    SetState(State)
    InState(State) bool

    // error of this session
    Error() error
    SetError(error)

    // last active time
    SetLastActiveTime(int)
    LastActiveTime() int

    Close() error
}
