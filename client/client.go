package client

import (
    "github.com/eaglc/lamer/router"
)

type Client interface {
    Init(...Option) error
    Options() Options
    Router() router.Router
    String() string
}

