package selector

import "github.com/eaglc/lamer/registry"

type Strategy interface {
    Do([] registry.Node) Next
}