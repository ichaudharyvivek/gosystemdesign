package generator

import "sync/atomic"

type IncrementalGenerator struct {
	current atomic.Int64
}

func NewIncrementalGenerator() *IncrementalGenerator {
	g := &IncrementalGenerator{}
	g.current.Store(1000000)
	return g
}

func (g *IncrementalGenerator) Generate() int64 {
	return g.current.Add(1)
}
