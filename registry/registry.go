package registry

type Registry interface {
    Init(...Option)
    Options() Options
    Register(Node) error
    Deregister(Node) error
    GetNodes(string) ([]Node, error)

    // TODO watcher
    //

    String() string
}

type Node interface {
    // TODO
}

type Marshaler interface {
    MarshalJNode() ([]byte, error)
}

type Unmarshaler interface {
    UnmarshalNode([] byte) error
}

func Marshal(v interface{}) ([]byte, error) {
    return nil,nil
}

func Unmarshal(data []byte, v interface{}) error  {
    return nil
}
