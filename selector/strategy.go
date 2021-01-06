package selector

import (
    "github.com/eaglc/lamer/registry"
)

type Strategy interface {
    Do(string, [] registry.Node) Next
    DoA(string, [] registry.Node) registry.Node
}