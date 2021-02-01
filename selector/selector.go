package selector

import "github.com/eaglc/lamer/registry"

type Selector interface {
    Init(...Option)
    Options() Options
    Select(...SelectOption) (Next, error)
    SelectA(...SelectOption) (registry.Node, error)

    // Mark a node error/success
    Mark(registry.Node, error)
    String() string
}

// return addr, error
type Next func() (registry.Node, error)

