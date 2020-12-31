package deliver

import "github.com/eaglc/lamer/session"

type Options struct {
    Cache session.Cache
}

type Option func(*Options)

func Cache(cache session.Cache) Option {
    return func(opts *Options) {
        opts.Cache = cache
    }
}