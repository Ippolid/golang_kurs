package storage

import (
	"math/rand/v2"
	"time"
)

func CleaningSession(r Storage, closeChan chan struct{}, n time.Duration) {
	for {
		select {
		case <-closeChan:
		case <-time.After(n):
			r.Clean()
		}
	}
}
func (r Storage) Clean() {
	k := rand.IntN(len(r.inner))
	k = k - k/2
	z := 0
	for i, _ := range r.inner {
		if r.inner[i].EXP <= time.Now().UnixMilli() && r.inner[i].EXP != 0 {
			delete(r.inner, i)
		}
		if z > k {
			break
		}
		z++
	}

}
