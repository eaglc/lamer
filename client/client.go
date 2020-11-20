package client

type Client interface {
    Init(...Option) error
    Options() Options
    String() string
}