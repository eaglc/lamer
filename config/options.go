package config

type Options struct {
	App string
}

type Option func(*Options)

func App(app string) Option {
	return func(o *Options) {
		o.App = app
	}
}
