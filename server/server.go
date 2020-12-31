package server

import (
    "github.com/eaglc/lamer/router"
)

type Server interface {
    Init(...Option)
    Options() Options
    Start() error
    Stop() error
    Router() router.Router
    String() string
}