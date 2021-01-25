package client

import (
    "context"
    "github.com/eaglc/lamer/codec"
    "github.com/eaglc/lamer/registry"
    "github.com/eaglc/lamer/router"
    "github.com/eaglc/lamer/selector"
    "github.com/eaglc/lamer/session"
    "github.com/eaglc/lamer/transport"
)

type Options struct {
    Addrs []string
    Codec     codec.NewCodec
    Registry  registry.Registry
    Selector  selector.Selector
    Transport transport.Transport
    Router    router.Router
    Callbacks []session.Callback
    Cache session.Cache

    Nodes []registry.Node

    DirectMode bool

    Name string
    Endpoints []string

    Context context.Context
}

type PoolOptions struct {

}

type Option func(*Options)

func Codec(codec codec.NewCodec) Option {
    return func(o *Options) {
        o.Codec = codec
    }
}

func Registry(r registry.Registry) Option {
    return func(o *Options) {
        o.Registry = r
    }
}

func Selector(s selector.Selector) Option {
    return func(o *Options) {
        o.Selector = s
    }
}

func Transport(t transport.Transport) Option {
    return func(o *Options) {
        o.Transport = t
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

func Addrs(addrs []string) Option {
    return func(o *Options) {
        o.Addrs = addrs
        o.DirectMode = true
    }
}

func Name(name string) Option {
    return func(o *Options) {
        o.Name = name
    }
}

func Endpoint(eds []string) Option {
    return func(o *Options) {
        o.Endpoints = eds
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

func Nodes(nodes ...registry.Node) Option {
    return func(o *Options) {
        o.Nodes = nodes
    }
}