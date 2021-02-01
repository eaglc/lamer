package selector

import (
    "github.com/eaglc/lamer/registry"
)

type Strategy interface {
    Do([]registry.Node, ...SelectOption) Next
    DoA([]registry.Node, ...SelectOption) registry.Node
    Mark(string, registry.Node, error)
    String() string
}