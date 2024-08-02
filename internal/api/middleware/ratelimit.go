package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"sync"
	"time"
)

type IPRateLimiter struct {
	rate     int64
	burst    int64
	m        sync.Map
	mu       sync.Mutex
	interval time.Duration
}

func NewIPRateLimiter(rate int64, interval time.Duration) *IPRateLimiter {
	return &IPRateLimiter{
		rate:     rate,
		burst:    rate * 10,
		interval: interval,
	}
}

func (i *IPRateLimiter) getBucket(key string) *ratelimit.Bucket {
	bucket, exists := i.m.Load(key)
	if !exists {
		i.mu.Lock()
		defer i.mu.Unlock()

		// 双重检查锁
		bucket, exists = i.m.Load(key)
		if !exists {
			bucket = ratelimit.NewBucketWithQuantum(i.interval, i.burst, i.rate)
			i.m.Store(key, bucket)
		}
	}
	return bucket.(*ratelimit.Bucket)
}

func (i *IPRateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.GetHeader("CF-Connecting-IP")
		if ip == "" {
			ip = c.GetHeader("X-Forwarded-For")
		}
		if ip == "" {
			ip = c.ClientIP()
		}
		bucket := i.getBucket(ip)

		if bucket.TakeAvailable(1) == 0 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}

		c.Next()
	}
}
