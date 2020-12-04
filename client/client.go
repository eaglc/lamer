package client

import "github.com/eaglc/lamer/router"

type Client interface {
    Init(...Option) error
    Options() Options
    String() string
    Router() router.Router
}