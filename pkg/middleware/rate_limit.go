package middleware

import (
	"net/http"
	"strings"
	"sync"
	"time"
)

type (
	Limiter struct {
		limit    int
		remains  int
		duration time.Duration
		limits   map[string]*limiter
		mu       sync.Mutex
	}
	limiter struct {
		limit     int
		remaining int
		resetTime int64
	}
)

func NewLimiter(limit int, remains int, duration time.Duration) *Limiter {
	return &Limiter{limit: limit, remains: remains, duration: duration, limits: make(map[string]*limiter)}
}

func (l *Limiter) RateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := strings.Split(r.RemoteAddr, ":")[0]
		l.mu.Lock()
		defer l.mu.Unlock()
		if l.limits[clientIP] == nil {
			l.limits[clientIP] = &limiter{
				limit:     l.limit,
				remaining: l.limit,
				resetTime: time.Now().Add(l.duration).Unix(),
			}
		}

		// if limit exists
		if l.limits[clientIP].remaining <= 0 {
			http.Error(w, "Too many requests, please try again later", http.StatusTooManyRequests)
			return
		}

		// check if rate limit reset is needed
		if time.Now().Unix() > l.limits[clientIP].resetTime {
			l.limits[clientIP].remaining = l.limits[clientIP].limit
			l.limits[clientIP].resetTime = time.Now().Add(l.duration).Unix()
		}

		l.limits[clientIP].remaining--

		next.ServeHTTP(w, r)
	}
}
