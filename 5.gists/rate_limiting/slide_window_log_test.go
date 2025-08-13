package rate_limiting

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSlidingWindowLog_AllowRequest(t *testing.T) {
	tests := []struct {
		name               string
		windowSize         int64
		maxRequests        int64
		requests           int
		shouldAllow        bool
		waitBetweenWindows bool
	}{
		{
			name:        "允许未超限请求",
			windowSize:  1,
			maxRequests: 3,
			requests:    2,
			shouldAllow: true,
		},
		{
			name:        "拒绝超限请求",
			windowSize:  1,
			maxRequests: 3,
			requests:    4,
			shouldAllow: false,
		},
		{
			name:               "窗口滑动后重置",
			windowSize:         1,
			maxRequests:        3,
			requests:           4,
			shouldAllow:        true,
			waitBetweenWindows: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swl := NewSlidingWindowLog(tt.windowSize, tt.maxRequests)

			var lastResult bool
			for i := 0; i < tt.requests; i++ {
				if tt.waitBetweenWindows && i == tt.requests-1 {
					time.Sleep(time.Duration(tt.windowSize+1) * time.Second)
				}
				lastResult = swl.AllowRequest()
			}

			if lastResult != tt.shouldAllow {
				t.Errorf("AllowRequest() = %v, want %v", lastResult, tt.shouldAllow)
			}
		})
	}
}

func TestSlidingWindowLog_ConcurrentAccess(t *testing.T) {
	swl := NewSlidingWindowLog(1, 1000)
	var wg sync.WaitGroup
	allowedCount := int64(0)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if swl.AllowRequest() {
				atomic.AddInt64(&allowedCount, 1)
			}
		}()
	}

	wg.Wait()

	if allowedCount != 1000 {
		t.Errorf("Expected 1000 allowed requests, got %d", allowedCount)
	}
}
