package server

import (
    "github.com/eaglc/lamer/router"
    "github.com/eaglc/lamer/transport"
)

type Server interface {
    Init(...Option)
    Options() Options
    Start() error
    Stop() error
    String() string
    // A server contains multiple transports, and each transport is bound to a router
    Router(t transport.Type) router.Router
}