package storage

import (
	"math/rand/v2"
	"time"
)

func CleaningSession(r *Storage, closeChan chan struct{}, n time.Duration) {
	for {
		select {
		case <-closeChan:
			return
		case <-time.After(n):
			r.Clean()
		}
	}
}
func (r *Storage) Clean() {
	k := rand.IntN(len(r.inner))
	k = k - k/2
	z := 0
	nowtime := time.Now().UnixMilli()
	for i := range r.inner {
		if r.inner[i].Exp <= nowtime && r.inner[i].Exp != 0 {
			delete(r.inner, i)
		}
		if z > k {
			break
		}
		z++
	}

}
