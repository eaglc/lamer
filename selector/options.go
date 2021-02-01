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

type SelectOptions struct {
    // strategy used
    Strategy string

    // endpoint name
    Name string

    // key for select
    Key string
}

type Option func(*Options)
type SelectOption func(*SelectOptions)
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

func StrategyName(s string) SelectOption {
    return func(o *SelectOptions) {
        o.Strategy = s
    }
}

func Endpoint(s string) SelectOption {
    return func(o *SelectOptions) {
        o.Name = s
    }
}

func Key(s string) SelectOption {
    return func(o *SelectOptions) {
        o.Key = s
    }
}
