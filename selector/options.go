package selector

import (
    "context"
    "github.com/eaglc/lamer/registry"
)

type Options struct {
    Getter registry.Getter
    Filters []Filter
    Strategies []Strategy

    ctx context.Context
}

type Option func(*Options)

func Strategies(s ...Strategy) Option {
    return func(o *Options) {
        o.Strategies = s
    }
}

func Filters(f ...Filter) Option {
    return func(o *Options) {
        o.Filters = f
    }
}

func Getter(r registry.Getter) Option {
    return func(o *Options) {
        o.Getter = r
    }
}