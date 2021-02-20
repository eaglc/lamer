package letter

import (
	"context"
)

type Request interface {
	Init(...RequestOption)
	Body() interface{}
	Options() RequestOptions
}

type RequestOptions struct {
	PeerAddr string
	PeerName string
	PeerId   string
	Context  context.Context
}

type RequestOption func(options *RequestOptions)

func PeerAddr(addr string) RequestOption {
	return func(o *RequestOptions) {
		o.PeerAddr = addr
	}
}

func PeerName(name string) RequestOption {
	return func(o *RequestOptions) {
		o.PeerName = name
	}
}

func PeerId(id string) RequestOption {
	return func(o *RequestOptions) {
		o.PeerId = id
	}
}

func RequestContext(ctx context.Context) RequestOption {
	return func(o *RequestOptions) {
		o.Context = ctx
	}
}
