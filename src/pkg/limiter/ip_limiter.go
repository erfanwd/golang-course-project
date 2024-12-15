package limiter

import (
	"sync"

	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
	return i
}

func (i *IPRateLimiter) AddIp(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter

	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
    i.mu.RLock() 
    limiter, exists := i.ips[ip]
    i.mu.RUnlock()

    if !exists {
        i.mu.Lock() 
        defer i.mu.Unlock()

        limiter, exists = i.ips[ip]
        if !exists {
            limiter = rate.NewLimiter(i.r, i.b)
            i.ips[ip] = limiter
        }
    }
    return limiter
}


