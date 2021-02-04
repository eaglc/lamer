package router

import (
    "context"
    "github.com/eaglc/lamer/handler"
)

// RouteBefore is executed on the request before the router.Serve
// One of the uses is to develop a distribution strategy
type RouteBefore func(ctx context.Context, req interface{}) context.Context

// ServeBefore is executed on the request before the handler
type ServeBefore func(ctx context.Context, req interface{}) context.Context
type ServeAfter func(ctx context.Context, replay interface{}) context.Context
type ServeFinalizer func(ctx context.Context, reply interface{}) error

type Router interface {
    handler.Handler
    Init(opts ...Option)
    Handle(pattern interface{}, handler handler.Func, opts ...Option)
    Start() error
    Stop() error
    String() string
}

//func Handle(pattern string, h handler.Handler)  {
//    DefaultRouter.Handle(pattern, h)
//}
//
//type router struct {
//
//}
//
//func (r *router)Handle(pattern interface{}, h handler.Handler)  {
//
//}
//
//var (
//    DefaultRouter = router{}
//)