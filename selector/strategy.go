package selector

import (
    "github.com/eaglc/lamer/registry"
)

type Strategy interface {
    Do(...SelectOption) Next
    DoA(...SelectOption) registry.Node
    Mark(string, registry.Node, error)
    String() string
}