package pkg

import "sync"

type device struct {
	*sync.Mutex
	ready bool
}

func (d *device) isReady() bool {
	return d.ready
}

func (d *device) processRequest(req *request) {
	d.Lock()
	d.ready = false
	d.Unlock()
}
