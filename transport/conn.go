package transport

type Conn interface {
    Init(...ConnOption)
    Close() error
    Read(b []byte) (n int, e error)
    Write(b []byte) (n int, e error)
    LocalAddr() string
    RemoteAddr() string
}