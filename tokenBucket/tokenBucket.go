package tokenBucket

import (
	"time"
)

type TokenBucket struct {
	// The maximum rate of the token bucket in tokens per second.
	maxBucketSize int64
	// The rate at which the token bucket refills in tokens per second.
	refillRate int64
	// The number of tokens currently in the bucket.
	currentTokens int64
	// The time at which the token bucket was last refilled.
	lastRefillTime time.Time
}

// NewTokenBucket creates a new TokenBucket with the given maxBucketSize and refillRate.
func NewTokenBucket(maxBucketSize int64, refillRate int64) *TokenBucket {
	return &TokenBucket{
		maxBucketSize:  maxBucketSize,
		refillRate:     refillRate,
		currentTokens:  maxBucketSize,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) Refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefillTime)
	tb.lastRefillTime = now
	tb.currentTokens += elapsed.Nanoseconds() * tb.refillRate / (1e+09)
	if tb.currentTokens > tb.maxBucketSize {
		tb.currentTokens = tb.maxBucketSize
	}
}

func (tb *TokenBucket) Consume(count int64) bool {
	tb.Refill()
	if tb.currentTokens >= count {
		tb.currentTokens -= count
		return true
	}
	return false
}
