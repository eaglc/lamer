package server

import (
    "github.com/eaglc/lamer/router"
)

type Server interface {
    Init(...Option)
    Options() Options
    Start() error
    Stop() error

    // A server contains multiple transports
    // A router is bound to one or more transports
    // A server can contain multiple routers
    Router(t string) router.Router
    DefaultRouter() router.Router

    String() string
}