package client

import (
    "github.com/eaglc/lamer/router"
)

type Client interface {
    Init(...Option)
    Options() Options
    Router() router.Router
    Name() string
    String() string
    Start() error
    Stop() error
}

