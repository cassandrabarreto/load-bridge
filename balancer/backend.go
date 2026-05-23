package balancer

import (
	"net/url"
	"sync"
)
type Backend struct {
	URL *url.URL
	alive bool
	mu sync.RWMutex
}

func NewBackend(rawUrl string) (*Backend, error){
	u , err := url.Parse(rawUrl)
	if err != nil{
		return nil, err
	}
	return &Backend{URL: u, alive: true} , nil
}

func (b *Backend) SetAlive (alive bool){
	b.mu.Lock()
	defer b.mu.Unlock()
	b.alive = alive
}

func (b *Backend) IsAlive() bool{
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.alive
}