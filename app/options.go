package app

import "context"

type Options struct {

    Context context.Context
}

type Option func(*Options)

func Context(ctx context.Context) Option {
    return func(o *Options) {
        o.Context = ctx
    }
}