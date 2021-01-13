package registry

import (
    "context"
)

type Getter interface {
    GetNodes(string) ([]Node, error)
}

type Registry interface {
    Getter
    Init(...Option)
    Options() Options
    Register(Node, ...RegisterOption) error
    Deregister(Node) error
    Watch(context.Context, string) (Watcher, error)
    String() string
}

type Node interface {
    Key() string
    Data() []byte
    Addr() string
}

type Marshaler func(Node) ([]byte, error)
type Unmarshaler func([]byte) (Node, error)