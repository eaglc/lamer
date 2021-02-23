package app

import (
	"github.com/eaglc/lamer/client"
	"github.com/eaglc/lamer/server"
	"os"
	"os/signal"
	"syscall"
)

type App interface {
	Init(...Option)
	Options() Options
	Name() string

	// a App can contains multiple clients
	Client(string) client.Client

	// a App contain only one server
	Server() server.Server

	Run() error
	String() string
}

type app struct {
	opts Options
}

func (a *app) Init(opts ...Option) {
	for _, o := range opts {
		o(&a.opts)
	}

	if a.opts.Server == nil {
		panic("Server is nil")
	}
}

func (a *app) Options() Options {
	return a.opts
}

func (a *app) Name() string {
	return a.opts.Server.Options().Name
}

func (a *app) Client(n string) client.Client {
	for i := range a.opts.Clients {
		if a.opts.Clients[i].Name() == n {
			return a.opts.Clients[i]
		}
	}

	return nil
}

func (a *app) Server() server.Server {
	return a.opts.Server
}

// To be optimized: start & stop
func (a *app) Run() error {
	// signal
	sch := make(chan os.Signal)
	signal.Notify(sch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// Must not block
	for i := range a.opts.Clients {
		if err := a.opts.Clients[i].Start(); err != nil {
			return err
		}
	}

	// Must not block
	if err := a.opts.Server.Start(); err != nil {
		return err
	}

	// block here to wait end
	select {
	case <-sch:
	case <-a.opts.Context.Done():
	}

	a.opts.Server.Stop()
	for i := range a.opts.Clients {
		a.opts.Clients[i].Stop()
	}

	return nil
}

func (a *app) String() string {
	return "lamer.app"
}

func NewApp(opts ...Option) App {
	var options Options
	for _, o := range opts {
		o(&options)
	}
	return &app{
		opts: options,
	}
}
