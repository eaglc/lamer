package app

import (
    "context"
    "github.com/eaglc/lamer/client"
    "github.com/eaglc/lamer/registry"
    "github.com/eaglc/lamer/router"
    "github.com/eaglc/lamer/server"
    "github.com/eaglc/lamer/transport"
)

type Options struct {

    ////
    // Sink into client/server, when client/server dose not have these options
    Registry registry.Registry
    Transport transport.Transport
    Router router.Router
    ////

    Server server.Server

    // To be optimized: map or array
    Clients []client.Client

    Context context.Context
}

type Option func(*Options)

func Context(ctx context.Context) Option {
    return func(o *Options) {
        o.Context = ctx
    }
}

func Registry(r registry.Registry) Option {
    return func(o *Options) {
        o.Registry = r
    }
}

func Server(s server.Server) Option {
    return func(o *Options) {
        o.Server = s
    }
}

func Client(c ...client.Client) Option {
    return func(o *Options) {
        o.Clients = c
    }
}

func Transport(t transport.Transport) Option {
    return func(o *Options) {
        o.Transport = t
    }
}

func Router(r router.Router) Option {
    return func(o *Options) {
        o.Router = r
    }
}