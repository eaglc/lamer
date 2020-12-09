package client

type Stream interface {
    SendMsg(m interface{}) error
    RecvMsg(m interface{}) error
}

