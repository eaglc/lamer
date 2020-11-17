package transport

import (
    "context"
    "crypto/tls"
    "time"
)

type Options struct {
    // dial timeout
    DialTimeout time.Duration

    // tls config
    TLSConfig *tls.Config

    // context
    Context context.Context
}

type Option func(*Options)

func DialTimeout(t time.Duration) Option {
    return func(o *Options) {
        o.DialTimeout = t
    }
}

func TLSConfig(tls *tls.Config) Option {
    return func(o *Options) {
        o.TLSConfig = tls
    }
}

func Context(ctx context.Context) Option {
    return func(o *Options) {
        o.Context = ctx
    }
}

////////////////////////////////////////////////////////////////////////////////
type ConnOptions struct {
    ReadDeadline time.Time
    WriteDeadline time.Time
}

type ConnOption func(*ConnOptions)

func ReadDeadline(t time.Time) ConnOption {
    return func(o *ConnOptions) {
        o.ReadDeadline = t
    }
}

func WriteDeadline(t time.Time) ConnOption {
    return func(o *ConnOptions) {
        o.WriteDeadline = t
    }
}