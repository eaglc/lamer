package client

type RawClient interface {
    Client
    Send(m interface{}) error
}