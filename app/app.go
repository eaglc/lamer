package app

import (
    "github.com/eaglc/lamer/client"
    "github.com/eaglc/lamer/server"
)

type App interface {
    Init(...Option)
    Options() Options
    Name() string
    Client() client.Client
    Server() server.Server
    Run() error
    String() string
}

type Handle interface {
    ServeConn()
}

func Handler(s string, handle Handle)  {

}