package transport

type Transport interface {
    Init(...Option) error
    Options() Options
    Dial(addr string) (Conn, error)
    Listen(addr string) (Listener, error)
    String() string
}

func NewTransport(opts ...Option) Transport {
    // TODO
    return nil
}