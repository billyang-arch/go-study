// 滑动窗口日志限流
package rate_limiting

import (
	"container/list"
	"sync"
	"time"
)

type SlidingWindowLog struct {
	windowSizeInSeconds int64
	maxRequestPerWindow int64
	requestLog          *list.List
	mu                  sync.Mutex
}

func NewSlidingWindowLog(windowSizeInSeconds, maxRequestPerWindow int64) *SlidingWindowLog {
	return &SlidingWindowLog{
		windowSizeInSeconds: windowSizeInSeconds,
		maxRequestPerWindow: maxRequestPerWindow,
		requestLog:          list.New(),
	}
}

func (swl *SlidingWindowLog) AllowRequest() bool {
	swl.mu.Lock()
	defer swl.mu.Unlock()

	now := time.Now().Unix()
	windowStart := now - swl.windowSizeInSeconds

	// remove all request log that is out of window
	for swl.requestLog.Len() > 0 && swl.requestLog.Front().Value.(int64) <= windowStart {
		swl.requestLog.Remove(swl.requestLog.Front())
	}

	if int64(swl.requestLog.Len()) < swl.maxRequestPerWindow {
		swl.requestLog.PushBack(now)
		return true
	}

	return false
}
