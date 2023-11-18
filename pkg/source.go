package pkg

import (
	"time"
)

type request struct {
	SourceID int
	GenTime  time.Time
}

func createRequest(num int) {
	for {
		if !start {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		in.ready(&request{
			SourceID: num,
			GenTime:  time.Now(),
		})
		if stop {
			time.Sleep(100 * time.Millisecond)
			return
		}
	}
}
