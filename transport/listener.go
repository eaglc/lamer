package transport

type Listener interface {
    Accept() (Conn, error)
    Close() error
    Addr() string
}