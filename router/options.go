package router

import "context"

type Options struct {
	RouteBefore RouteBefore
	Before      []ServeBefore
	After       []ServeAfter
	Finalizer   []ServeFinalizer
	PoolSize    int
	ChanSize    int
	ChanCount   int
	Context     context.Context
}

type Option func(*Options)

func Route(b RouteBefore) Option {
	return func(o *Options) {
		o.RouteBefore = b
	}
}

func Before(b ...ServeBefore) Option {
	return func(o *Options) {
		o.Before = b
	}
}

func After(a ...ServeAfter) Option {
	return func(o *Options) {
		o.After = a
	}
}

func Final(f ...ServeFinalizer) Option {
	return func(o *Options) {
		o.Finalizer = f
	}
}

func PoolSize(s int) Option {
	return func(o *Options) {
		o.PoolSize = s
	}
}

func ChanSize(s int) Option {
	return func(o *Options) {
		o.ChanSize = s
	}
}

func ChanCount(c int) Option {
	return func(o *Options) {
		o.ChanCount = c
	}
}

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}
