package registry

type Registry interface {
    Init(...Option)
    Options() Options
    Register(Node) error
    Deregister(Node) error
    GetNodes(Key) ([]Node, error)

    // TODO watcher

    String() string
}

type Node interface {
    Key
    Data
}

type Key interface {
    MarshalKey() ([]byte, error)
    UnmarshalKey() ([]byte, error)
}

type Data interface {
    MarshalData() ([]byte, error)
    UnmarshalData() ([]byte, error)
}