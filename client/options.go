package client

import (
    "context"
    "github.com/eaglc/lamer/codec"
    "github.com/eaglc/lamer/registry"
    "github.com/eaglc/lamer/router"
    "github.com/eaglc/lamer/selector"
    "github.com/eaglc/lamer/transport"
)

type Options struct {
    Codec     codec.NewCodec
    Registry  registry.Registry
    Selector  selector.Selector
    Transport transport.Transport
    Router    router.Router

    Context context.Context
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