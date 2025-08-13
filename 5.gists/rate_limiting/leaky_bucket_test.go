package rate_limiting

import (
	"sync"
	"testing"
	"time"
)

func TestLeakyBucket_AllowRequest(t *testing.T) {
	lb := NewLeakyBucket(5, 1.0) // 容量5，每秒漏出1个

	// 测试未超限请求
	for i := 0; i < 5; i++ {
		if !lb.AllowRequest() {
			t.Errorf("Expected request %d to be allowed", i)
		}

	}

	// 测试超限请求
	if lb.AllowRequest() {
		t.Error("Expected request to be denied when bucket is full")
	}

	// 测试漏出后重置
	time.Sleep(1 * time.Second)
	if !lb.AllowRequest() {
		t.Error("Expected request to be allowed after leak")
	}
}

func TestLeakyBucket_ConcurrentAccess(t *testing.T) {
	lb := NewLeakyBucket(1000, 10.0)
	var wg sync.WaitGroup

	// 模拟1000个并发请求
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lb.AllowRequest()
		}()
	}
	wg.Wait()

	// 验证最终计数
	if count := lb.CurrentCount(); count != 1000 {
		t.Errorf("Expected count 1000, got %d", count)
	}
}

func TestLeakyBucket_EdgeCases(t *testing.T) {
	// 测试零容量桶
	zeroBucket := NewLeakyBucket(0, 1.0)
	if zeroBucket.AllowRequest() {
		t.Error("Expected zero-capacity bucket to always deny requests")
	}

	// 测试极高漏出速率
	fastBucket := NewLeakyBucket(5, 1000.0)
	fastBucket.AllowRequest()
	time.Sleep(2 * time.Second)
	fastBucket.leak()
	if count := fastBucket.CurrentCount(); count != 0 {
		t.Errorf("Expected empty bucket after high leak rate, got %d", count)
	}
}
