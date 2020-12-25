package registry

import (
    "context"
)

type Getter interface {
    GetNodes(string) ([]Node, error)
}

type Watcher interface {
    Create(Node)
    Update(Node)
    Delete(Node)
}

type NewWatcher func(context.Context) Watcher

type Registry interface {
    Getter
    Init(...Option)
    Options() Options
    Register(Node, ...RegisterOption) error
    Deregister(Node) error
    // TODO watcher
    // You can have different watcher to handle event for different key
    Watch(context.Context, string, NewWatcher) error

    String() string
}

type Node interface {
    MarshalKey() ([]byte, error)
    MarshalNode() ([]byte, error)
    UnmarshalNode([]byte) error

    // param: transport
    Addrs(string) []string
}