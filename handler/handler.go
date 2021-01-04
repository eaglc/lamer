package handler

import (
    "context"
)

type Handler interface {
    Serve(ctx context.Context, request interface{}, replay interface{}) error
}

type Func func(ctx context.Context, request interface{}, replay interface{}) error

func (f Func)Serve(ctx context.Context, request interface{}, replay interface{}) error {
    return f(ctx, request, replay)
}
