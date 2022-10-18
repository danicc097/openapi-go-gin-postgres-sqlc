// https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
// TODO check improvements from above:
// Check the X-Forwarded-For or X-Real-IP headers for the IP address, if you are running your server behind a reverse proxy.
// Switch to a sync.RWMutex to help reduce contention on the map.

package rest

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// Create a custom visitor struct which holds the rate limiter for each
// visitor and the last time that the visitor was seen.
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// rateLimitMiddleware allows rate limiting requests.
type rateLimitMiddleware struct {
	Logger *zap.Logger
	// rlim is the number of events per second allowed.
	rlim     rate.Limit
	visitors map[string]*visitor

	mu sync.Mutex
}

func newRateLimitMiddleware(
	logger *zap.Logger,
	rlim rate.Limit,
) *rateLimitMiddleware {
	return &rateLimitMiddleware{
		Logger:   logger,
		rlim:     rlim,
		visitors: make(map[string]*visitor),
		mu:       sync.Mutex{},
	}
}

// Limit is the middleware function to rate limits requests.
func (r *rateLimitMiddleware) Limit() gin.HandlerFunc {
	go r.cleanupVisitors(3 * time.Minute)

	return func(c *gin.Context) {
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			log.Print(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		limiter := r.getVisitor(ip)
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)

			return
		}

		c.Next()
	}
}

func (r *rateLimitMiddleware) getVisitor(ip string) *rate.Limiter {
	r.mu.Lock()
	defer r.mu.Unlock()

	v, exists := r.visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(r.rlim, 3)
		// Include the current time when creating a new visitor.
		r.visitors[ip] = &visitor{limiter, time.Now()}

		return limiter
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()

	return v.limiter
}

// Every minute check the map for visitors that haven't been seen for
// more than the given duration and delete the entries.
func (r *rateLimitMiddleware) cleanupVisitors(d time.Duration) {
	for {
		time.Sleep(time.Minute)

		r.mu.Lock()
		for ip, v := range r.visitors {
			if time.Since(v.lastSeen) > d {
				delete(r.visitors, ip)
			}
		}
		r.mu.Unlock()
	}
}
