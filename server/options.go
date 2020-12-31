package server

import (
    "context"
    "github.com/eaglc/lamer/codec"
    "github.com/eaglc/lamer/registry"
    "github.com/eaglc/lamer/router"
    "github.com/eaglc/lamer/session"
    "github.com/eaglc/lamer/transport"
)

type Options struct {
    // session callback
    Callbacks []session.Callback

    // session cache
    Cache session.Cache

    // Registry
    Registry registry.Registry

    // Transport
    Transport transport.Transport

    // Router
    Router router.Router

    // Codec
    Codec codec.NewCodec

    // Address to listen
    Addr string

    // server name
    Name string

    // server id
    Id string

    // server version
    Version string

    // context
    Context context.Context
}

type Option func(*Options)

func Codec(c codec.NewCodec) Option {
    return func(o *Options) {
        o.Codec = c
    }
}

func Registry(r registry.Registry) Option {
    return func(o *Options) {
        o.Registry = r
    }
}

func Transport(t transport.Transport) Option {
    return func(o *Options) {
        o.Transport = t
    }
}

func Name(name string) Option {
    return func(o *Options) {
        o.Name = name
    }
}

func Id(id string) Option {
    return func(o *Options) {
        o.Id = id
    }
}

func Version(v string) Option {
    return func(o *Options) {
        o.Version = v
    }
}

func Context(ctx context.Context) Option {
    return func(o *Options) {
        o.Context = ctx
    }
}

func Router(r router.Router) Option {
    return func(o *Options) {
        o.Router = r
    }
}

func Addr(addr string) Option {
    return func(o *Options) {
        o.Addr = addr
    }
}

func Callbacks(cbs ...session.Callback) Option {
    return func(o *Options) {
        o.Callbacks = cbs
    }
}

func Cache(cache session.Cache) Option {
    return func(o *Options) {
        o.Cache = cache
    }
}