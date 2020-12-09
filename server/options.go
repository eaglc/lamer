package server

import (
    "context"
    "github.com/eaglc/lamer/codec"
    "github.com/eaglc/lamer/registry"
    "github.com/eaglc/lamer/router"
    "github.com/eaglc/lamer/transport"
)

type Options struct {
    // registry
    Registry registry.Registry

    // we can have multiple transports
    Transports map[string] transport.Transport

    // default Router
    Router router.Router

    // Router map
    BindRouter map[string]router.Router

    // default Codec
    Codec codec.NewCodec

    // Codec map
    BindCodec map[string] codec.NewCodec

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

func Transport(t ...transport.Transport) Option {
    return func(o *Options) {
        if o.Transports == nil {
            o.Transports = make(map[string] transport.Transport)
        }
        for i := range t {
            o.Transports[t[i].String()] = t[i]
        }
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

func BindRouter(t transport.Transport, r router.Router) Option {
    return func(o *Options) {
        if o.BindRouter == nil {
            o.BindRouter = make(map[string]router.Router)
        }
        if o.Transports == nil {
            o.Transports = make(map[string] transport.Transport)
        }

        o.BindRouter[t.String()] = r
        o.Transports[t.String()] = t
    }
}

func BindCodec(t transport.Transport, c codec.NewCodec) Option {
    return func(o *Options) {
        if o.BindCodec == nil {
            o.BindCodec = make(map[string] codec.NewCodec)
        }
        if o.Transports == nil {
            o.Transports = make(map[string] transport.Transport)
        }

        o.BindCodec[t.String()] = c
        o.Transports[t.String()] = t
    }
}

func BindRouterCodec(t transport.Transport, r router.Router, c codec.NewCodec) Option {
    return func(o *Options) {
        if o.BindCodec == nil {
            o.BindCodec = make(map[string] codec.NewCodec)
        }
        if o.Transports == nil {
            o.Transports = make(map[string] transport.Transport)
        }
        if o.BindRouter == nil {
            o.BindRouter = make(map[string]router.Router)
        }

        o.Transports[t.String()] = t
        o.BindRouter[t.String()] = r
        o.BindCodec[t.String()] = c
    }
}