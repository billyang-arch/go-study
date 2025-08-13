package rate_limiting

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTokenBucket_AllowRequest(t *testing.T) {
	tests := []struct {
		name         string
		capacity     int64
		fillRate     float64
		requestToken int64
		wantAllow    bool
		sleep        time.Duration
	}{
		{
			name:         "允许未超限请求",
			capacity:     10,
			fillRate:     1,
			requestToken: 5,
			wantAllow:    true,
		},
		{
			name:         "拒绝超限请求",
			capacity:     10,
			fillRate:     1,
			requestToken: 15,
			wantAllow:    false,
		},
		{
			name:         "填充后允许请求",
			capacity:     10,
			fillRate:     1,
			requestToken: 5,
			wantAllow:    true,
			sleep:        6 * time.Second, // 填充6个token
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := NewTokenBucket(tt.capacity, tt.fillRate)

			if tt.sleep > 0 {
				time.Sleep(tt.sleep)
			}

			got := tb.AllowRequest(tt.requestToken)
			if got != tt.wantAllow {
				t.Errorf("AllowRequest() = %v, want %v", got, tt.wantAllow)
			}
		})
	}
}

func TestTokenBucket_ConcurrentAccess(t *testing.T) {
	tb := NewTokenBucket(1000, 10)
	var wg sync.WaitGroup
	allowed := int64(0)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if tb.AllowRequest(1) {
				atomic.AddInt64(&allowed, 1)
			}
		}()
	}

	wg.Wait()

	if allowed != 1000 {
		t.Errorf("并发测试失败: 允许 %d 次请求, 预期 1000", allowed)
	}
}
