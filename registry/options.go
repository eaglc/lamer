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
    Encode Marshaler
    Decode Unmarshaler

    Context context.Context
}

type RegisterOptions struct {
    TTL time.Duration
}

type Option func(*Options)

type RegisterOption func(* RegisterOptions)

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

func RegisterTTL(ttl time.Duration) RegisterOption {
    return func(o *RegisterOptions) {
        o.TTL = ttl
    }
}

func Context(ctx context.Context) Option {
    return func(o *Options) {
        o.Context = ctx
    }
}

func Encoder(marshal Marshaler) Option {
    return func(o *Options) {
        o.Encode = marshal
    }
}

func Decoder(unmarshal Unmarshaler) Option {
    return func(o *Options) {
        o.Decode = unmarshal
    }
}