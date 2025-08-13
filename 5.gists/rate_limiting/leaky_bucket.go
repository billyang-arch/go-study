package rate_limiting

import (
	"sync"
	"sync/atomic"
	"time"
)

type LeakyBucket struct {
	capacity          int64
	leakRate          float64
	count             int64 // atomic counter
	lastLeakTimestamp time.Time
	mu                sync.Mutex
}

func NewLeakyBucket(capacity int64, leakRate float64) *LeakyBucket {
	if leakRate <= 0 {
		panic("leakRate must be positive")
	}
	return &LeakyBucket{
		capacity:          capacity,
		leakRate:          leakRate,
		lastLeakTimestamp: time.Now(),
	}
}

func (lb *LeakyBucket) AllowRequest() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.leak()

	current := atomic.LoadInt64(&lb.count)
	if current < lb.capacity {
		atomic.AddInt64(&lb.count, 1)
		return true
	}
	return false
}

func (lb *LeakyBucket) leak() {
	now := time.Now()
	elapsed := now.Sub(lb.lastLeakTimestamp)
	if elapsed <= 0 {
		return
	}

	leakAmount := int64(elapsed.Seconds() * lb.leakRate)
	if leakAmount > 0 {
		current := atomic.LoadInt64(&lb.count)
		newCount := current - leakAmount
		if newCount < 0 {
			newCount = 0
		}
		atomic.StoreInt64(&lb.count, newCount)
		lb.lastLeakTimestamp = now
	}
}

// 新增监控方法
func (lb *LeakyBucket) CurrentCount() int64 {
	return atomic.LoadInt64(&lb.count)
}
