package selector

import "context"

type Options struct {
    Strategies []Strategy

    ctx context.Context
}

type Option func(*Options)

func Strategies(s ...Strategy) Option {
    return func(o *Options) {
        o.Strategies = s
    }
}