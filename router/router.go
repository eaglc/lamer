package router

import "github.com/eaglc/lamer/handler"

type Router interface {
    Handle(pattern interface{}, handler handler.Handler)
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