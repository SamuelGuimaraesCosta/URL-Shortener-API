package middleware

import (
	"net/http"
	"sync"
	"time"
)

type visitor struct {
	lastSeen time.Time
	count    int
}

var visitors = make(map[string]*visitor)
var mu sync.Mutex

const limit = 100

func RateLimit(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr

		mu.Lock()

		v, exists := visitors[ip]

		if !exists {
			visitors[ip] = &visitor{count: 1, lastSeen: time.Now()}
		} else {

			if v.count > limit {
				mu.Unlock()
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			v.count++
			v.lastSeen = time.Now()
		}

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
