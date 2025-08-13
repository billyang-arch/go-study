package rate_limiting

import (
	"sync"
	"time"
)

type FixedWindowCounter struct {
	windowSizeInSeconds int64 // Size of the window in seconds
	maxRequestPerWindow int64 // Maximum number of requests allowed per window
	currentWindowStart  int64 // Start time of the current window (epoch seconds)
	requestCount        int64 // Number of requests in the current window
	mu                  sync.Mutex
}

func NewFixedWindowCounter(windowSizeInSeconds, maxRequestPerWindow int64) *FixedWindowCounter {
	return &FixedWindowCounter{
		windowSizeInSeconds: windowSizeInSeconds,
		maxRequestPerWindow: maxRequestPerWindow,
		currentWindowStart:  time.Now().Unix(),
		requestCount:        0,
	}
}

func (fwc *FixedWindowCounter) AllowRequest() bool {
	// 加锁，防止并发问题
	fwc.mu.Lock()
	defer fwc.mu.Unlock()

	now := time.Now().Unix()

	// Check if we 've moved to a new window
	if now > fwc.currentWindowStart+fwc.windowSizeInSeconds {
		fwc.currentWindowStart = now
		fwc.requestCount = 0
	}

	// Check if we've reached the maximum number of requests
	if fwc.requestCount >= fwc.maxRequestPerWindow {
		return false
	}

	// Increment the request count and return true
	fwc.requestCount++
	return true
}
