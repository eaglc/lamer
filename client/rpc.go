package client

import "context"

type RpcClient interface {
    Client
    Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error
    Stream(ctx context.Context, serviceMethod string, args interface{}) (Stream, error)
}