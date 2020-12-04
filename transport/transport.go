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

type Type int32

const (
    TCP Type = iota
    UDP
    HTTP1
    HTTP2
    WEBSOCKET
    QUIC
    RPC
    GRPC
)

func (tt Type) String() string {
    switch tt {
    case TCP:
        return "tcp"
    case UDP:
        return "udp"
    case HTTP1:
        return "http1"
    case HTTP2:
        return "http2"
    case WEBSOCKET:
        return "ws"
    case QUIC:
        return "quic"
    case RPC:
        return "rpc"
    case GRPC:
        return "grpc"
    default:
        return "lamer"
    }
}