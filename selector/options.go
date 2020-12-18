package selector

import (
    "context"
    "github.com/eaglc/lamer/registry"
)

type Options struct {
    Registry registry.Registry
    Strategies []Strategy

    ctx context.Context
}

type Option func(*Options)

func Strategies(s ...Strategy) Option {
    return func(o *Options) {
        o.Strategies = s
    }
}

func Registry(r registry.Registry) Option {
    return func(o *Options) {
        o.Registry = r
    }
}