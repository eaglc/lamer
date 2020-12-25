package selector

import "github.com/eaglc/lamer/registry"

type Filter interface {
    Do([]registry.Node) []registry.Node
}