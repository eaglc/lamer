package registry

import (
    "context"
)

type Watcher interface {
    Create(Node)
    Update(Node)
    Delete(Node)
}

type NewWatcher func(context.Context) Watcher

type Registry interface {
    Init(...Option)
    Options() Options
    Register(Node) error
    Deregister(Node) error
    GetNodes(string) ([]Node, error)

    // TODO watcher
    // You can have different watcher to handle event for different key
    Watch(context.Context, string, NewWatcher) error

    String() string
}

type Node interface {
    MarshalKey() ([]byte, error)
    MarshalNode() ([]byte, error)
    UnmarshalNode([]byte) error
    Addr() string
}