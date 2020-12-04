package handler

import "context"

type Handler interface {
    Handle(ctx context.Context, request interface{}, replay interface{}) error
}

type Func func(ctx context.Context, request interface{}, replay interface{}) error

func (f Func)Handle(ctx context.Context, request interface{}, replay interface{}) error {
    return f(ctx, request, replay)
}
