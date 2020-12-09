package registry

import (
    "context"
    "crypto/tls"
    "time"
)

type Options struct {
    Addrs []string
    Timeout time.Duration
    TLSConfig *tls.Config

    Context context.Context
}

type Option func(*Options)

func Addrs(addrs ...string) Option {
    return func(o *Options) {
        o.Addrs = addrs
    }
}

func Timeout(timeout time.Duration) Option {
    return func(o *Options) {
        o.Timeout = timeout
    }
}

func TLSConfig(t *tls.Config) Option {
    return func(o *Options) {
        o.TLSConfig = t
    }
}