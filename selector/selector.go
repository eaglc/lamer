package selector

import "github.com/eaglc/lamer/registry"

type Selector interface {
    Init(...Option)
    Options() Options
    Select(string) (Next, error)

    // Mark a node error/success
    Mark(string, node registry.Node, err error)
    String() string
}

// return addr, error
type Next func() (registry.Node, error)

