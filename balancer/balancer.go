package balancer

import (
	"sync/atomic"
)

type Balancer struct {
	backends []*Backend
	counter atomic.Uint64
}

func NewBalancer(backends []*Backend) *Balancer {
	return &Balancer{backends: backends}
}

func (b *Balancer) NextBackend() *Backend {
	total := uint64(len(b.backends))

	for range total{
		backendIndex := b.counter.Add(1) % total
		backend := b.backends[backendIndex]
		if backend.IsAlive(){
			return backend
		}
	}
	return nil
}