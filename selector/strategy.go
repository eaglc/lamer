package selector

import (
    "github.com/eaglc/lamer/registry"
)

type Strategy interface {
    Do(string, ...SelectOption) Next
    DoA(string, ...SelectOption) registry.Node
    Mark(registry.Node, error)
    String() string
}