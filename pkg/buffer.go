package pkg

import "sync"

type buffer struct {
	sync.Mutex
	buf         []*request
	iterIn      int
	iterDecline int // newest request in buffer
}

func (b *buffer) getHandler() (int, error) {
	b.Lock()
	defer b.Unlock()
	return b.iterIn, nil
}

func (b *buffer) put(req *request) {
	b.Lock()
	defer b.Unlock()
}

func (b *buffer) get() (*request, error) {
	b.Lock()
	defer b.Unlock()
	return nil, nil
}

func (b *buffer) remove(handler int) *request {
	return nil
}
