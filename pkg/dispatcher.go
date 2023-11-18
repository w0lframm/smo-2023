package pkg

import (
	"sync"
	"time"
)

type inDIspatcher struct {
	sync.Mutex
}

type outDispatcher struct {
	devices []device
	iter    int
}

func (s *inDIspatcher) ready(req *request) {
	s.Lock()
	defer s.Unlock()
	handler, err := b.getHandler()
	if err != nil {
		_ = b.remove(handler)
		// report decline of reqToDecline
	}
	b.put(req)
}

func (out *outDispatcher) run() {
	for {
		for _, device := range out.devices {
			if device.isReady() {
				req, err := b.get()
				for err == nil {
					time.Sleep(100 * time.Millisecond)
					req, err = b.get()
				}
				device.processRequest(req)
			}
		}
		if stop {
			time.Sleep(100 * time.Millisecond)
			return
		}
	}
}
