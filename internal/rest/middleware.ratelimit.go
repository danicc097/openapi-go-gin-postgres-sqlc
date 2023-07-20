// https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
// TODO check improvements from above:
// Check the X-Forwarded-For or X-Real-IP headers for the IP address, if you are running your server behind a reverse proxy.
// Switch to a sync.RWMutex to help reduce contention on the map.

package rest

import (
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
	logger *zap.SugaredLogger
	// rlim is the number of events per second allowed.
	rlim rate.Limit
	// blim is the number of burst allowed.
	blim     int
	visitors map[string]*visitor // don't mind pointers, it's for internal struct use only

	mu sync.Mutex
}

func newRateLimitMiddleware(
	logger *zap.SugaredLogger, rlim rate.Limit,
	blim int,
) *rateLimitMiddleware {
	return &rateLimitMiddleware{
		logger:   logger,
		rlim:     rlim,
		blim:     blim,
		visitors: make(map[string]*visitor),
		mu:       sync.Mutex{},
	}
}

// Limit is the middleware function to rate limits requests.
func (m *rateLimitMiddleware) Limit() gin.HandlerFunc {
	go m.cleanupVisitors(3 * time.Minute)

	return func(c *gin.Context) {
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			ip = c.Request.RemoteAddr
			if ip == "" {
				ip = "unknown"
			}
		}
		m.logger.Infof("ip: %v", ip)

		limiter := m.getVisitor(ip)
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)

			return
		}

		c.Next()
	}
}

func (m *rateLimitMiddleware) getVisitor(ip string) *rate.Limiter {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, exists := m.visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(m.rlim, m.blim)
		// Include the current time when creating a new visitor.
		m.visitors[ip] = &visitor{limiter, time.Now()}

		return limiter
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()

	return v.limiter
}

// Every minute check the map for visitors that haven't been seen for
// more than the given duration and delete the entries.
func (m *rateLimitMiddleware) cleanupVisitors(d time.Duration) {
	for {
		time.Sleep(time.Minute)

		m.mu.Lock()
		for ip, v := range m.visitors {
			if time.Since(v.lastSeen) > d {
				delete(m.visitors, ip)
			}
		}
		m.mu.Unlock()
	}
}
