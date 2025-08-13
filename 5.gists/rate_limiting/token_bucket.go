package rate_limiting

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity            int64      // Maximum number of tokens the bucket can hold
	fillRate            float64    // Rate at which tokens are added to the bucket (tokens per second)
	tokens              int64      // Current number of tokens in the bucket
	lastRefillTimestamp time.Time  // Last time we refilled the bucket
	mu                  sync.Mutex // Mutex for thread safety
}

func NewTokenBucket(capacity int64, fillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:            capacity,
		fillRate:            fillRate,
		tokens:              capacity, // Start with a full bucket
		lastRefillTimestamp: time.Now(),
	}
}

func (tb *TokenBucket) AllowRequest(tokens int64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill() // First, add any new tokens based on elapsed time

	if tb.tokens < tokens {
		return false // Not enough tokens, deny the request
	}

	tb.tokens -= tokens // Consume the tokens
	return true         // Allow the request
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefillTimestamp)
	tokensToAdd := tb.fillRate * elapsed.Seconds()

	tb.tokens = tb.tokens + int64(tokensToAdd)
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefillTimestamp = now
}
