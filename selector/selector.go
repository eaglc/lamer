package selector

import "github.com/eaglc/lamer/registry"

type Selector interface {
    Init(...Option)
    Options() Options
    Select(string) Next
    String() string
}

type Next func() (registry.Node, error)

