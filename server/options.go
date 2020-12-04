package server

import (
    "context"
    "github.com/eaglc/lamer/codec"
    "github.com/eaglc/lamer/registry"
    "github.com/eaglc/lamer/transport"
)

type Options struct {
    Codec codec.NewCodec
    Registry registry.Registry
    Transport []transport.Transport
    Name string
    Id string
    Version string

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